import type {Settings as LayoutSettings} from '@ant-design/pro-layout';
import {PageLoading} from '@ant-design/pro-layout';
import type {RunTimeLayoutConfig} from 'umi';
import {history, Link} from 'umi';
import RightContent from '@/components/RightContent';
import Footer from '@/components/Footer';
import {BookOutlined, LinkOutlined} from '@ant-design/icons';
import {getCurrentUser} from "@/services/user";
import type {API} from "@/services/user/typing";
import {message} from 'antd';
import type {RequestConfig} from "@@/plugin-request/request";
import {ErrorShowType} from "@@/plugin-request/request";
import {oauthLogin} from "@/services/login";


const isDev = process.env.NODE_ENV === 'development';
const loginPath = '/user/login';
const registryPath = '/user/register';


/** 获取用户信息比较慢的时候会展示一个 loading */
export const initialStateConfig = {
  loading: <PageLoading/>,
};

/**
 * @see  https://umijs.org/zh-CN/plugins/plugin-initial-state
 * 启用 umijs/plugin-initial-state：有 src/app.ts 并且导出 getInitialState 方法时启用
 * 功能：会在整个应用最开始执行，返回值会作为全局共享的数据，并通过 useModel('@@initialState') 直接获取数据
 *      也可以用来导出全局函数
 * */
export async function getInitialState(): Promise<{
  settings?: Partial<LayoutSettings>; // ProLayout配置
  currentUser?: API.CurrentUser;  // 当前的登录用户，如果不存在将会跳转到登录页面
  fetchUserInfo?: () => Promise<API.CurrentUser | undefined>; // 获取登录的用户信息函数
}> {
  // 获取当前用户信息
  const fetchUserInfo = async () => {
    // 判断是否有code参数，如果有代表第三方oauth登录
    const ps = new URLSearchParams(history.location.search);
    const code = ps.get("code");
    const state = ps.get("state");
    if (code) { // 有code，进行github 登录
      const resp = await oauthLogin({
        code: code,
        redirect_url: window.location.href,
        type: "github",
        state: state || ""
      });
      if (resp.code == 200) {
        // 设置token
        localStorage.setItem("Token", resp.data.token);
      }
    }

    if (!localStorage.getItem("Token")) {
      history.push(loginPath);
      return undefined;
    }
    try {
      const resp = await getCurrentUser();
      if (resp.code == 200) {
        return resp.data;
      }
      if (resp.code == 420) { // token失效
        // 跳转登录页面
        history.push(loginPath);
        return undefined;
      }
      history.push(loginPath);
      // 弹出错误提示
      message.error(resp.msg)
      return undefined
    } catch (error) {
      history.push(loginPath);
    }
    return undefined;
  };

  // 如果是登录或者注册页面，不执行
  if (history.location.pathname !== loginPath
    && history.location.pathname !== registryPath) {
    const currentUser = await fetchUserInfo();
    return {
      settings: {},
      fetchUserInfo,
      currentUser: currentUser,
    };
  }
  return {
    settings: {},
    fetchUserInfo,
  };
}


/**
 * 请求request对象全局封装配置
 * @see https://umijs.org/zh-CN/plugins/plugin-request
 * @see https://github.com/umijs/umi-request/blob/master/README_zh-CN.md
 * umijs/plugin-request
 * 基于 umi-request 和 ahooks 的 useRequest 提供了一套统一的网络请求和错误处理方案。
 */
export const request: RequestConfig = {
  // 请求url前缀
  prefix: '/api/v1',
  // 超时时间
  timeout: 300 * 1000,
  // 公共请求头
  headers: {
    'Accept': 'application/json',
    'Content-Type': 'application/json',
    'Authorization': 'Bearer ' + localStorage.getItem('Token')
  },
  errorConfig: {
    // 当后端返回的数据类型不符合umi定义的接口规范时，使用adaptor进行适配
    adaptor: (resData) => {
      return {
        ...resData,
        // 交易是否成功
        success: resData.code == 200,
        // 交易成功返回的数据
        data: resData.data,
        // 交易失败的错误码
        errorCode: resData.code,
        // 交易错误的信息, 这里可以通过adaptor的context对象或者resData对象获取状态码
        // 如果http状态码为200 那么resData.status字段不存在，反之存在
        errorMessage: resData.msg ? resData.msg : "网络通讯失败",
        // 错误的显示类型
        // 0 silent; 1 message.warn; 2 message.error; 4 notification; 9 page redirect
        // 0 无错误提示  1 警告信息     3 错误信息        4 通知           9 页面跳转
        showType: ErrorShowType.ERROR_MESSAGE
      }
    }
  },
  middlewares: [],
  requestInterceptors: [],
  responseInterceptors: [],
};

// ProLayout 支持的api https://procomponents.ant.design/components/layout
// ProLayout 可以提供一个标准又不失灵活的中后台标准布局，同时提供一键切换布局形态，
// 自动生成菜单等功能。与 PageContainer 配合使用可以自动生成面包屑，页面标题，并且提供低成本方案接入页脚工具栏。
export const layout: RunTimeLayoutConfig = ({initialState}) => {
  return {
    // 自定义头右部的 render 方法
    rightContentRender: () => <RightContent/>,
    disableContentMargin: false,
    // 关闭面包屑渲染
    breadcrumbRender: false,
    // 关闭页面标题渲染
    // pageTitleRender: false,
    // 配置水印，水印是 PageContainer 的功能，layout 只是透传给 PageContainer
    // waterMarkProps: {
    //   content: initialState?.currentUser?.userid,
    // },
    // 自定义页脚
    footerRender: () => <Footer/>,
    // 页面切换时触发，这里判断用户是否登录，如果没有登录，跳转到登录页面
    onPageChange: () => {
      const {location} = history;
      // 如果没有登录，重定向到 login
      if (!initialState?.currentUser
        && location.pathname !== loginPath
        && history.location.pathname !== registryPath) {
        history.push(loginPath);
      }
    },
    // 显示在菜单右下角的快捷操作
    links: isDev
      ? [
        <Link to="/umi/plugin/openapi" target="_blank">
          <LinkOutlined/>
          <span>OpenAPI 文档</span>
        </Link>,
        <Link to="/~docs">
          <BookOutlined/>
          <span>业务组件文档</span>
        </Link>,
      ]
      : [],
    // 自定义的菜单头区域
    menuHeaderRender: undefined,
    // 自定义 403 页面
    // unAccessible: <div>unAccessible</div>,
    ...initialState?.settings,
  };


};
