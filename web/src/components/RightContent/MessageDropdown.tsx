import React, {useCallback} from 'react';
import {BellOutlined} from '@ant-design/icons';
import {Menu, Spin} from 'antd';
import {history, useModel} from 'umi';
import HeaderDropdown from '../HeaderDropdown';
import styles from './index.less';
import type {MenuInfo} from 'rc-menu/lib/interface';

export type GlobalHeaderRightProps = {
  menu?: boolean;
};

const MessageDropdown: React.FC<GlobalHeaderRightProps> = ({}) => {
  const {initialState} = useModel('@@initialState');

  const onMenuClick = useCallback(
    (event: MenuInfo) => {
      const {key} = event;
      if (key === 'system') {
        console.log("点击了系统消息");
      }
      history.push(`/account/${key}`);
    },
    [],
  );

  const loading = (
    <span className={`${styles.action} ${styles.account}`}>
      <Spin
        size="small"
        style={{
          marginLeft: 8,
          marginRight: 8,
        }}
      />
    </span>
  );

  if (!initialState) {
    return loading;
  }

  const {currentUser} = initialState;

  if (!currentUser || !currentUser.nickname) {
    return loading;
  }

  const menuHeaderDropdown = (
    <Menu className={styles.menu} selectedKeys={[]} onClick={onMenuClick}>
      <Menu.Item key="system">
        系统通知
      </Menu.Item>
      <Menu.Item key="stars">
        评论关注
      </Menu.Item>
      <Menu.Item key="setting">
        消息设置
      </Menu.Item>
    </Menu>
  );
  return (
    <HeaderDropdown overlay={menuHeaderDropdown}>
      <span className={`${styles.action} ${styles.account}`}>
        <BellOutlined/>
      </span>
    </HeaderDropdown>
  );
};

export default MessageDropdown;
