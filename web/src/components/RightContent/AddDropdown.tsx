import React, {useCallback} from 'react';
import {CaretDownFilled, PlusOutlined} from '@ant-design/icons';
import {Menu, Spin} from 'antd';
import {history, useModel} from 'umi';
import {stringify} from 'querystring';
import HeaderDropdown from '../HeaderDropdown';
import styles from './index.less';
import {outLogin} from '@/services/ant-design-pro/api';
import type {MenuInfo} from 'rc-menu/lib/interface';

export type GlobalHeaderRightProps = {
  menu?: boolean;
};

/**
 * 退出登录，并且将当前的 url 保存
 */
const loginOut = async () => {
  await outLogin();
  const {query = {}, pathname} = history.location;
  const {redirect} = query;
  // Note: There may be security issues, please note
  if (window.location.pathname !== '/user/login' && !redirect) {
    history.replace({
      pathname: '/user/login',
      search: stringify({
        redirect: pathname,
      }),
    });
  }
};

const AddDropdown: React.FC<GlobalHeaderRightProps> = ({}) => {
  const {initialState, setInitialState} = useModel('@@initialState');

  const onMenuClick = useCallback(
    (event: MenuInfo) => {
      const {key} = event;
      if (key === 'logout') {
        setInitialState((s) => ({...s, currentUser: undefined}));
        loginOut();
        return;
      }
      history.push(`/space/${key}`);
    },
    [setInitialState],
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
      <Menu.Item key={"new"}>
        创建空间
      </Menu.Item>
      <Menu.Item>
        创建组织
      </Menu.Item>
    </Menu>
  );
  return (
    <HeaderDropdown overlay={menuHeaderDropdown}>
      <span className={`${styles.action} ${styles.account}`}>
        <PlusOutlined style={{marginRight: 2}}/>
        <CaretDownFilled style={{fontSize: 11}}/>
      </span>
    </HeaderDropdown>
  );
};

export default AddDropdown;
