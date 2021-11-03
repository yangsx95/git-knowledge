import React, {useEffect, useState} from 'react';
import {PageContainer} from '@ant-design/pro-layout';
import {Card, Cascader} from 'antd';
import ProForm, {
  ProFormText,
  ProFormSelect,
  ProFormList,
  ProFormGroup,
  ProFormDependency,
} from '@ant-design/pro-form';
import {getCredentials, getGitOrgs, getGitRepos, getOrganizations} from "@/services/user";
import ProFormItem from "@ant-design/pro-form/es/components/FormItem";
import type {CascaderOptionType} from "antd/lib/cascader";

export default (): React.ReactNode => {

  // 页面状态（是否是第一次加载）
  const [onload, setOnload] = useState<boolean>(true);
  // 凭据列表
  const [credentials, setCredentials] = useState<{ value: string, label: string, }[]>([]);

  // 请求用户所有组织
  const requestOrganizations = () => {
    return async () => {
      // 查询用户加入的所有组织
      const os = await getOrganizations();
      const showOrganizations: { value: string, label: string }[] = [];
      os.data.forEach(o => {
        showOrganizations.push({label: `${o.org_id}`, value: o.org_id,})
      });
      return showOrganizations
    };
  }

  // 获取指定凭据下的所有git组织
  const getGitOrganizations = (credentialId: string): CascaderOptionType[] => {
    const ops: CascaderOptionType[] = [];
    if (!credentialId) {
      return ops;
    }
    (async () => {
      const resp = await getGitOrgs(credentialId);
      if (resp.code == 200) {
        resp.data.forEach(e => {
          ops.push({label: e.org_id, value: e.org_id, isLeaf: false})
        });
      }
    })()
    return ops;
  }

  const loadData = (credentialId: string) => {
    return async (selectedOptions?: CascaderOptionType[]) => {
      if (!selectedOptions) {
        return
      }
      if (!credentialId) {
        return
      }
      const targetOption = selectedOptions[selectedOptions.length - 1];
      targetOption.loading = true;
      // 加载组织下的所有仓库
      const repos = await getGitRepos(credentialId, selectedOptions[0].value as string);
      // 添加子节点
      const arr = new Array<CascaderOptionType>();
      repos.data.forEach(e => {
        arr.push({
          label: e.name,
          value: e.name,
        })
      })
      targetOption.children = arr;
      targetOption.loading = false;
    };
  };

  useEffect(() => {
    // 页面初始化数据
    if (onload) {
      setOnload(false);
      // 加载凭据列表
      (async () => {
        const cs = await getCredentials();
        const scs: { value: string, label: string, }[] = [];
        cs.data.forEach(o => {
          scs.push({label: `${o.name}`, value: o.credential_id,})
        });
        setCredentials(scs);
      })();
    }
    return () => {
    }
  }, [onload]);

  return (
    // PageContainer 封装了 antd 的 PageHeader 组件，增加了 tabList 和 content。
    // 根据当前的路由填入 title 和 breadcrumb。它依赖 Layout 的 route 属性。当然你可以传入参数来复写默认值。
    // PageContainer 支持 Tabs 和 PageHeader 的所有属性。
    <PageContainer title={"创建一个空间"} subTitle={"空间可以包含多个git仓库，用于存储您知识库的数据"}>
      <Card>
        <ProForm onFinish={async (values) => console.log(values)}>
          <ProForm.Group label="基本信息">
            <ProFormSelect label="所属" name="owner" request={requestOrganizations()} width={120} required={true}/>
            <ProFormText label="空间名称" name="name" tooltip="最长为 8 位" placeholder="请输入空间名称" width={300} required={true}/>
            <ProFormText label="描述" name="description" width={400}/>
          </ProForm.Group>
          <ProForm.Group label="主仓库">
            <ProFormSelect label="API凭据" name="credential_id" options={credentials} width={120} required={true}/>
            <ProFormDependency name={["credential_id"]} shouldUpdate={true}>
              {(({credential_id}) => {
                return (
                  <ProFormItem label="选择仓库" tooltip="配置仓库，定义了space空间的配置" style={{width: 300}} required={true}
                               shouldUpdate={true}  name={"main_repository_id"}>
                    <Cascader options={getGitOrganizations(credential_id)} loadData={loadData(credential_id)}/>
                  </ProFormItem>
                );
              })}
            </ProFormDependency>
          </ProForm.Group>
          <ProForm.Group label="子仓库">
            <ProFormList
              name="child_repositories"
              rules={[
                {
                  validator: async (_, value) => {
                    if (value && value.length > 0) {
                      return;
                    }
                    throw new Error('至少要有一项！');
                  },
                },
              ]}
              creatorRecord={{
                name: '',
              }}
              initialValue={[]}
            >
              <ProFormGroup>
                <ProFormSelect label="API凭据" name="credential_id" options={credentials} width={120} required={true}/>
                <ProFormDependency name={["credential_id"]} shouldUpdate={true}>
                  {(({credential_id}) => {
                    return (
                      <ProFormItem label="选择仓库" tooltip="配置仓库，定义了space空间的配置" style={{width: 300}} required={true}
                                   name={"repository_id"} shouldUpdate={true}>
                        <Cascader options={getGitOrganizations(credential_id)} loadData={loadData(credential_id)}/>
                      </ProFormItem>
                    );
                  })}
                </ProFormDependency>
                <ProFormText label="仓库名称" name="repository_name" tooltip="最长为 8 位" width={300} required={true}/>
              </ProFormGroup>
            </ProFormList>
          </ProForm.Group>
        </ProForm>
      </Card>
    </PageContainer>
  );
};
