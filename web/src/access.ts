import type {API} from "@/services/user/typing";

/**
 * @see https://umijs.org/zh-CN/plugins/plugin-access
 * umijs/plugin-access
 * 启用：有 src/access.ts 时启用。
 * 约定了 src/access.ts 为我们的权限定义文件，该文件需要默认导出一个方法，导出的方法会在项目初始化时被执行。
 * 该方法需要返回一个对象，对象的每一个值就对应定义了一条权限
 * */
export default function access(initialState: { currentUser?: API.CurrentUser | undefined }) {
  const { currentUser } = initialState || {};
  return {
    canAdmin: currentUser,
  };
}
