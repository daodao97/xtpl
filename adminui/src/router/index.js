import { Layout } from '@okiss/oms'

export const routes = [
  {
    path: '/',
    component: Layout,
    meta: {
      menuType: 1
    },
    children: [
      {
          path: '/example',
          name: 'Exmaple',
          meta: { title: 'Example', menuType: 1, keepAlive: false },
          role: 'root',
          children: [
            // {
            //   path: '/example/sub',
            //   name: 'ExmapleSub',
            //   component: () => import('../views/example/index.vue'),
            //   meta: { title: 'ExampleSub', icon: 'el-icon-data-analysis', menuType: 2, keepAlive: false },
            //   role: 'root'
            // }
          ]
        }
    ]
  }
]
