export default [
  {
    path: '/user',
    layout: false,
    routes: [
      {
        path: '/user',
        routes: [
          {
            name: 'login',
            path: '/user/login',
            component: './User/Login',
          },
          {
            name: 'register',
            path: '/user/register',
            component: './User/Register',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },

  {
    path: '/space',
    routes: [
      {
        path: '/space',
        routes: [
          {
            path: '/space/:orgId/:spaceId',
            component: './Space',
          },
          {
            name: 'new',
            path: '/space/new',
            component: './Space/New',
          },
        ],
      },
      {
        component: './404',
      },
    ],
  },

  {
    path: '/welcome',
    // name属性决定该菜单是否展示
    // name: 'welcome',
    icon: 'smile',
    component: './Welcome',
  },


  {
    path: '/',
    redirect: '/welcome',
  },
  {
    component: './404',
  },
];
