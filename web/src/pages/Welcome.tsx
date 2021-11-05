import React, {useEffect, useState} from 'react';
import {Avatar, Card, Divider, Layout} from "antd";
import Sider from "antd/es/layout/Sider";
import {Content} from "antd/es/layout/layout";
import {DownOutlined} from "@ant-design/icons";
import ProList from '@ant-design/pro-list';
import {useModel} from "@@/plugin-model/useModel";
import {getAllSpaces} from "@/services/space";

export default (): React.ReactNode => {
  const {initialState} = useModel('@@initialState');
  // @ts-ignore
  const {currentUser} = initialState;
  const [onload, setOnload] = useState<boolean>(true);
  const [listData, setListData] = useState<{ name: string; image: string; }[]>([]);

  useEffect(() => {
    if (onload) {
      setOnload(false);
      (async () => {
        const ss = await getAllSpaces();
        const showSpaces: { name: string; image: string; }[] = [];
        ss.data.forEach(o => {
          showSpaces.push({name: o.name, image: currentUser.avatar_url})
        });
        setListData(showSpaces);
      })();
    }
    return () => {
    }
  }, [currentUser.avatar_url, onload]);

  return (<>
    <Layout>
      <Sider style={{backgroundColor: "white"}} width={"25%"}>
        <Card bordered={false}>
          <Avatar size={"small"} icon={<Avatar src={currentUser.avatar_url}/>}/> {currentUser.nickname}
          <DownOutlined style={{fontSize: 11}}/>
        </Card>
        <Divider dashed style={{marginTop: 0, marginBottom: 0}}/>
        <ProList
          onRow={(record: any) => {
            return {
              onMouseEnter: () => {
                console.log(record);
              },
              onClick: () => {
                console.log(record);
              },
            };
          }}
          headerTitle={"Spaces"}
          rowKey="name"
          metas={{
            title: {
              dataIndex: 'name',
            },
            avatar: {
              dataIndex: 'image',
            },
          }}
          dataSource={listData}
        />
      </Sider>
      <Content style={{backgroundColor: "white"}}>Content</Content>
    </Layout>
  </>);
};
