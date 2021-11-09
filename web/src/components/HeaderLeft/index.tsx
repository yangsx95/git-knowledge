import {Avatar, Button, Drawer, List, Space} from 'antd';
import React, {useEffect, useState} from 'react';
import {useModel} from 'umi';
import {getAllSpaces} from "@/services/space";
import {DownOutlined, MenuUnfoldOutlined, WechatOutlined} from "@ant-design/icons";
import Settings from "../../../config/defaultSettings";
import {Link} from "@umijs/preset-dumi/lib/theme";


const HeaderLeft: React.FC = () => {
  const {initialState} = useModel('@@initialState');
  // @ts-ignore
  const {currentUser} = initialState;
  const [onload, setOnload] = useState<boolean>(true);
  const [listData, setListData] = useState<{ name: string; image: string; }[]>([]);
  const [visible, setVisible] = useState(false);
  const showDrawer = () => {
    setVisible(true);
  };
  const onClose = () => {
    setVisible(false);
  };

  // https://blog.csdn.net/pig_html/article/details/114838699
  // 不要在循环，条件或嵌套函数中调用 Hook， 确保总是在你的 React 函数的最顶层以及任何 return 之前调用他们
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


  return (
    <>
      <Space>
        <Button icon={<MenuUnfoldOutlined/>} style={{border: "none", marginLeft: 5}} onClick={showDrawer}/>
        {Settings.title}
      </Space>
      <Drawer
        title={<Space><Avatar size={"small"} src={currentUser.avatar_url}/>{currentUser.nickname}<DownOutlined
          style={{fontSize: 11}}/></Space>}
        placement={"left"}
        width={300}
        onClose={onClose}
        visible={visible}
      >
        <List
          itemLayout="horizontal"
          dataSource={listData}
          renderItem={item => (
            <List.Item>
              <List.Item.Meta
                title={<Link to={`/space/${currentUser.userid}/${item.name}`} onClick={onClose} style={{fontSize: 12}}>
                  <Space><WechatOutlined/>{item.name}</Space>
                </Link>}/>
            </List.Item>
          )}
        />
      </Drawer>
    </>
  );
};
export default HeaderLeft;
