import React, {useState} from 'react';
import {PageContainer} from '@ant-design/pro-layout';
import {Card, Cascader} from 'antd';
import ProForm, {
  ProFormText,
  ProFormSelect,
  ProFormList,
  ProFormGroup,
} from '@ant-design/pro-form';
import {getCredentials, getGitOrgs, getGitRepos, getOrganizations} from "@/services/user";
import ProFormItem from "@ant-design/pro-form/es/components/FormItem";
import type {CascaderOptionType} from "antd/lib/cascader";

export default (): React.ReactNode => {

  const [currentCredId, setCurrentCredId] = useState<string>("");

  const [gitOrgs, setGitOrgs] = useState<{
    value: string,
    label: string,
    isLeaf: boolean,
  }[]>([]);

  function getOrgs() {
    return async () => {
      // 查询用户加入的所有组织
      const os = await getOrganizations();
      const showOrgs: {
        value: string,
        label: string,
      }[] = [];
      os.data.forEach(o => {
        showOrgs.push({
          label: `${o.org_id}`,
          value: o.org_id,
        })
      });
      return showOrgs
    };
  }

  function getCreds() {
    return async () => {
      const cs = await getCredentials();
      const showCreds: {
        value: string,
        label: string,
      }[] = [];
      cs.data.forEach(o => {
        showCreds.push({
          label: `${o.name}`,
          value: o.credential_id,
        })
      });
      return showCreds
    };
  }


  const loadData = async (selectedOptions?: CascaderOptionType[]) => {
    if (!selectedOptions) {
      return
    }

    const targetOption = selectedOptions[selectedOptions.length - 1];
    targetOption.loading = true;

    // 加载组织下的所有仓库
    const repos = await getGitRepos(currentCredId, selectedOptions[0].value as string);
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

  // 更改凭证触发的事件（获取凭证下的所有组织）
  const changeCred = async (value: string) => {
    setCurrentCredId(value)
    const go = await getGitOrgs(value);
    if (go.code == 200) {
      const ops: {
        label: string
        value: string
        isLeaf: boolean
      }[] = [];
      go.data.forEach(e => {
        ops.push({label: e.org_id, value: e.org_id, isLeaf: false})
      });
      setGitOrgs(ops);
    }
  }

  return (
    // PageContainer 封装了 antd 的 PageHeader 组件，增加了 tabList 和 content。
    // 根据当前的路由填入 title 和 breadcrumb。它依赖 Layout 的 route 属性。当然你可以传入参数来复写默认值。 PageContainer 支持 Tabs 和 PageHeader 的所有属性。
    <PageContainer title={"创建一个空间"} subTitle={"空间可以包含多个git仓库，用于存储您知识库的数据"}>
      <Card style={{width: '70%', marginLeft: '15%'}}>
        <ProForm
          onFinish={async (values) => console.log(values)}>
          <ProForm.Group label="基本信息">
            <ProFormSelect
              request={getOrgs()}
              name="owner"
              label="所属"
              width={120}
              required={true}
            />
            <ProFormText
              name="name"
              label="空间名称"
              tooltip="最长为 8 位"
              placeholder="请输入空间名称"
              width={300}
              required={true}
            />
            <ProFormText name={"description"} label="描述" width={400}/>
          </ProForm.Group>
          <ProForm.Group label="主仓库">
            <ProFormSelect
              request={getCreds()}
              width={120}
              name="credential_id"
              label="API凭据"
              fieldProps={{onChange: changeCred}}
            />
            <ProFormItem label="选择仓库" name="repositoryUrl" tooltip="配置仓库，定义了space空间的配置" style={{width: 300}}>
              <Cascader options={gitOrgs} loadData={loadData}/>
            </ProFormItem>
          </ProForm.Group>
          <ProForm.Group label="子仓库">
            <ProFormList
              name="users"
              rules={[
                {
                  validator: async (_, value) => {
                    console.log(value);
                    if (value && value.length > 0) {
                      return;
                    }
                    throw new Error('至少要有一项！');
                  },
                },
              ]}
              creatorRecord={{
                name: 'test',
              }}
              initialValue={[
                {
                  name: '1111',
                  nickName: '1111',
                  age: 111,
                  birth: '2021-02-18',
                  sex: 'man',
                  addrList: [{addr: ['taiyuan', 'changfeng']}],
                },
              ]}
            >
              <ProFormGroup>
                <ProFormSelect
                  options={[
                    {
                      value: 'fasd',
                      label: 'gitlab',
                    },
                    {
                      value: 'fa',
                      label: 'github',
                    },
                    {
                      value: 'gia',
                      label: 'gitlab-私服',
                    },
                  ]}
                  width={120}
                  name="main-repository"
                  label="凭据"
                />
                <ProFormSelect
                  width={300}
                  options={[
                    {
                      value: 'time',
                      label: 'notes',
                    },
                    {
                      value: 'tes',
                      label: 'spring',
                    },
                  ]}
                  tooltip="配置仓库，定义了space空间的配置"
                  name="repositoryUrl"
                  label="选择仓库"
                />
                <ProFormText
                  name="name"
                  label="仓库名称"
                  tooltip="最长为 8 位"
                  width={300}
                />
              </ProFormGroup>
            </ProFormList>
          </ProForm.Group>
        </ProForm>
      </Card>
    </PageContainer>
  );
};
