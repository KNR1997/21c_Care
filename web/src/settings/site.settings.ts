import {
  adminAndOwnerOnly,
  adminOnly,
  adminOwnerAndStaffOnly,
  ownerAndStaffOnly,
} from '@/utils/auth-utils';
import { Routes } from '@/config/routes';

export const siteSettings = {
  name: 'PickBazar',
  description: '',
  logo: {
    url: '/logo.svg',
    alt: 'PickBazar',
    href: '/',
    width: 138,
    height: 34,
  },
  collapseLogo: {
    url: '/collapse-logo.svg',
    alt: 'P',
    href: '/',
    width: 32,
    height: 32,
  },
  defaultLanguage: 'en',
  author: {
    name: 'RedQ',
    websiteUrl: 'https://redq.io',
    address: '',
  },
  headerLinks: [],
  authorizedLinks: [
    {
      href: Routes.profileUpdate,
      labelTransKey: 'authorized-nav-item-profile',
      icon: 'UserIcon',
      permission: adminOwnerAndStaffOnly,
    },
    // {
    //   href: Routes.shop.create,
    //   labelTransKey: 'common:text-create-shop',
    //   icon: 'ShopIcon',
    //   permission: adminAndOwnerOnly,
    // },

    // {
    //   href: Routes.settings,
    //   labelTransKey: 'authorized-nav-item-settings',
    //   icon: 'SettingsIcon',
    //   permission: adminOnly,
    // },
    {
      href: Routes.logout,
      labelTransKey: 'authorized-nav-item-logout',
      icon: 'LogOutIcon',
      permission: adminOwnerAndStaffOnly,
    },
  ],
  currencyCode: 'USD',
  sidebarLinks: {
    admin: {
      root: {
        href: Routes.dashboard,
        label: 'Main',
        icon: 'DashboardIcon',
        childMenu: [
          {
            href: Routes.dashboard,
            label: 'sidebar-nav-item-dashboard',
            icon: 'DashboardIcon',
          },
        ],
      },

      user: {
        href: '',
        label: 'text-user-control',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: '',
            label: 'text-patients',
            icon: 'UsersIcon',
            childMenu: [
              {
                href: Routes.patient.list,
                label: 'text-all-patients',
                icon: 'UsersIcon',
              },
              {
                href: Routes.patient.create,
                label: 'text-new-patients',
                icon: 'UsersIcon',
              },
            ],
          },
          {
            href: '',
            label: 'text-visits',
            icon: 'ChatIcon',
            childMenu: [
              {
                href: Routes.visit.list,
                label: 'text-all-visits',
                icon: 'UsersIcon',
              },
              {
                href: Routes.visit.create,
                label: 'text-new-visits',
                icon: 'UsersIcon',
              },
            ],
          },
        ],
      },

      catalog: {
        href: '',
        label: 'text-all-catalog',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: '',
            label: 'text-drugs',
            icon: 'MedicinePillsIcon',
            childMenu: [
              {
                href: Routes.drug.list,
                label: 'text-all-drugs',
                icon: 'MedicinePillsIcon',
              },
              {
                href: Routes.drug.create,
                label: 'text-new-drug',
                icon: 'MedicinePillsIcon',
              },
            ],
          },
          {
            href: '',
            label: 'text-lab-tests',
            icon: 'FlaskLabMedicalIcon',
            childMenu: [
              {
                href: Routes.labTest.list,
                label: 'text-all-lab-tests',
                icon: 'FlaskLabMedicalIcon',
              },
              {
                href: Routes.labTest.create,
                label: 'text-new-lab-tests',
                icon: 'FlaskLabMedicalIcon',
              },
            ],
          },
        ],
      },

      billing: {
        href: '',
        label: 'text-all-billings',
        icon: 'BillIcon',
        childMenu: [
          {
            href: '',
            label: 'text-billings',
            icon: 'BillIcon',
            childMenu: [
              {
                href: Routes.billing.list,
                label: 'text-all-billings',
                icon: 'BillIcon',
              },
            ],
          },
        ],
      },

      // settings: {
      //   href: '',
      //   label: 'text-site-management',
      //   icon: 'SettingsIcon',
      //   childMenu: [
      //     {
      //       href: Routes.settings,
      //       label: 'sidebar-nav-item-settings',
      //       icon: 'SettingsIcon',
      //       childMenu: [
      //         {
      //           href: Routes.settings,
      //           label: 'text-general-settings',
      //           icon: 'SettingsIcon',
      //         },
      //         {
      //           href: Routes.paymentSettings,
      //           label: 'text-payment-settings',
      //           icon: 'RefundsIcon',
      //         },
      //         {
      //           href: Routes.seoSettings,
      //           label: 'text-seo-settings',
      //           icon: 'StoreNoticeIcon',
      //         },
      //         {
      //           href: Routes.eventSettings,
      //           label: 'text-events-settings',
      //           icon: 'RefundsIcon',
      //         },
      //         {
      //           href: Routes.shopSettings,
      //           label: 'text-shop-settings',
      //           icon: 'RefundsIcon',
      //         },
      //         {
      //           href: Routes?.maintenance,
      //           label: 'text-maintenance-settings',
      //           icon: 'InformationIcon',
      //         },
      //         {
      //           href: Routes?.companyInformation,
      //           label: 'text-company-settings',
      //           icon: 'InformationIcon',
      //         },
      //         {
      //           href: Routes?.promotionPopup,
      //           label: 'text-popup-settings',
      //           icon: 'InformationIcon',
      //         },
      //         // {
      //         //   href: '',
      //         //   label: 'Social settings',
      //         //   icon: 'RefundsIcon',
      //         // },
      //       ],
      //     },
      //     // {
      //     //   href: '',
      //     //   label: 'Company Information',
      //     //   icon: 'InformationIcon',
      //     // },
      //     // {
      //     //   href: '',
      //     //   label: 'Maintenance',
      //     //   icon: 'MaintenanceIcon',
      //     // },
      //   ],
      // },
    },

    shop: {
      root: {
        href: '',
        label: 'text-main',
        icon: 'DashboardIcon',
        childMenu: [
          {
            href: (shop: string) => `${Routes.dashboard}${shop}`,
            label: 'sidebar-nav-item-dashboard',
            icon: 'DashboardIcon',
            permissions: adminOwnerAndStaffOnly,
          },
        ],
      },

      // analytics: {
      //   href: (shop: string) => `/${shop}${Routes.product.list}`,
      //   label: 'Analytics',
      //   icon: 'ShopIcon',
      //   permissions: adminAndOwnerOnly,
      //   childMenu: [
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Shop',
      //       icon: 'ShopIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Product',
      //       icon: 'ProductsIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Order',
      //       icon: 'OrdersIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Sale',
      //       icon: 'ShopIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //   ],
      // },

      product: {
        href: '',
        label: 'text-product-management',
        icon: 'ProductsIcon',
        permissions: adminOwnerAndStaffOnly,
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.product.list}`,
            label: 'sidebar-nav-item-products',
            icon: 'ProductsIcon',
            childMenu: [
              {
                href: (shop: string) => `/${shop}${Routes.product.list}`,
                label: 'text-all-products',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) => `/${shop}${Routes.product.create}`,
                label: 'text-new-products',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) => `/${shop}${Routes.draftProducts}`,
                label: 'text-my-draft',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) =>
                  `/${shop}${Routes.outOfStockOrLowProducts}`,
                label: 'text-all-out-of-stock',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
            ],
          },
          {
            href: (shop: string) => `/${shop}${Routes.productInventory}`,
            label: 'text-inventory',
            icon: 'InventoryIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.attribute.list}`,
            label: 'sidebar-nav-item-attributes',
            icon: 'AttributeIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.manufacturer.list}`,
            label: 'sidebar-nav-item-manufacturers',
            icon: 'DiaryIcon',
            permissions: adminAndOwnerOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.author.list}`,
            label: 'sidebar-nav-item-authors',
            icon: 'FountainPenIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },

      financial: {
        href: '',
        label: 'text-financial-management',
        icon: 'WithdrawIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.withdraw.list}`,
            label: 'sidebar-nav-item-withdraws',
            icon: 'AttributeIcon',
            permissions: adminAndOwnerOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.refund.list}`,
            label: 'sidebar-nav-item-refunds',
            icon: 'RefundsIcon',
            permissions: adminOwnerAndStaffOnly,
          },
        ],
      },

      order: {
        href: '',
        label: 'text-order-management',
        icon: 'OrdersIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.order.list}`,
            label: 'sidebar-nav-item-orders',
            icon: 'OrdersIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.transaction}`,
            label: 'text-transactions',
            icon: 'CalendarScheduleIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },

      feature: {
        href: '',
        label: 'text-feature-management',
        icon: 'ProductsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.storeNotice.list}`,
            label: 'sidebar-nav-item-store-notice',
            icon: 'StoreNoticeIcon',
            permissions: adminAndOwnerOnly,
          },
          {
            href: (shop: string) => `${Routes.ownerDashboardMessage}`,
            label: 'sidebar-nav-item-message',
            icon: 'ChatIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },

      feedback: {
        href: '',
        label: 'text-feedback-control',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.reviews.list}`,
            label: 'sidebar-nav-item-reviews',
            icon: 'ReviewIcon',
            permissions: adminAndOwnerOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.question.list}`,
            label: 'sidebar-nav-item-questions',
            icon: 'QuestionIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },

      user: {
        href: '',
        label: 'text-user-control',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.staff.list}`,
            label: 'sidebar-nav-item-staffs',
            icon: 'UsersIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },

      promotional: {
        href: '',
        label: 'text-promotional-control',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.coupon.list}`,
            label: 'Coupons',
            icon: 'CouponsIcon',
            permissions: adminAndOwnerOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.flashSale.list}`,
            label: 'text-flash-sale',
            icon: 'UsersIcon',
            childMenu: [
              {
                href: (shop: string) => `/${shop}${Routes.flashSale.list}`,
                label: 'text-available-flash-deals',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) =>
                  `/${shop}${Routes.myProductsInFlashSale}`,
                label: 'text-my-products-in-deals',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) =>
                  `/${shop}${Routes.vendorRequestForFlashSale.list}`,
                label: 'Ask for enrollment',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
            ],
          },
        ],
      },

      layout: {
        href: '',
        label: 'text-page-management',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.faqs.list}`,
            label: 'text-faqs',
            icon: 'TypesIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.termsAndCondition.list}`,
            label: 'Terms And Conditions',
            icon: 'TypesIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },
    },

    staff: {
      root: {
        href: '',
        label: 'text-main',
        icon: 'DashboardIcon',
        childMenu: [
          {
            href: (shop: string) => `${Routes.dashboard}${shop}`,
            label: 'sidebar-nav-item-dashboard',
            icon: 'DashboardIcon',
            permissions: adminOwnerAndStaffOnly,
          },
        ],
      },

      // analytics: {
      //   href: (shop: string) => `/${shop}${Routes.product.list}`,
      //   label: 'Analytics',
      //   icon: 'ShopIcon',
      //   permissions: adminAndOwnerOnly,
      //   childMenu: [
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Shop',
      //       icon: 'ShopIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Product',
      //       icon: 'ProductsIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Order',
      //       icon: 'OrdersIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.product.list}`,
      //       label: 'Sale',
      //       icon: 'ShopIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //   ],
      // },

      product: {
        href: '',
        label: 'text-product-management',
        icon: 'ProductsIcon',
        permissions: adminOwnerAndStaffOnly,
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.product.list}`,
            label: 'sidebar-nav-item-products',
            icon: 'ProductsIcon',
            childMenu: [
              {
                href: (shop: string) => `/${shop}${Routes.product.list}`,
                label: 'text-all-products',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) => `/${shop}${Routes.product.create}`,
                label: 'text-new-products',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) => `/${shop}${Routes.draftProducts}`,
                label: 'text-my-draft',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) =>
                  `/${shop}${Routes.outOfStockOrLowProducts}`,
                label: 'text-low-out-of-stock',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
            ],
          },
          {
            href: (shop: string) => `/${shop}${Routes.productInventory}`,
            label: 'text-inventory',
            icon: 'InventoryIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.attribute.list}`,
            label: 'sidebar-nav-item-attributes',
            icon: 'AttributeIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.manufacturer.list}`,
            label: 'sidebar-nav-item-manufacturers',
            icon: 'DiaryIcon',
            permissions: adminAndOwnerOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.author.list}`,
            label: 'sidebar-nav-item-authors',
            icon: 'FountainPenIcon',
            permissions: adminAndOwnerOnly,
          },
        ],
      },

      financial: {
        href: '',
        label: 'text-financial-management',
        icon: 'WithdrawIcon',
        childMenu: [
          // {
          //   href: (shop: string) => `/${shop}${Routes.withdraw.list}`,
          //   label: 'sidebar-nav-item-withdraws',
          //   icon: 'AttributeIcon',
          //   permissions: adminAndOwnerOnly,
          // },
          {
            href: (shop: string) => `/${shop}${Routes.refund.list}`,
            label: 'sidebar-nav-item-refunds',
            icon: 'RefundsIcon',
            permissions: adminOwnerAndStaffOnly,
          },
        ],
      },

      order: {
        href: '',
        label: 'text-order-management',
        icon: 'OrdersIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.order.list}`,
            label: 'sidebar-nav-item-orders',
            icon: 'OrdersIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          // {
          //   href: (shop: string) => `/${shop}${Routes.transaction}`,
          //   label: 'Transactions',
          //   icon: 'CalendarScheduleIcon',
          //   permissions: adminAndOwnerOnly,
          // },
        ],
      },

      // feature: {
      //   href: '',
      //   label: 'Features Management',
      //   icon: 'ProductsIcon',
      //   childMenu: [
      //     {
      //       href: (shop: string) => `/${shop}${Routes.storeNotice.list}`,
      //       label: 'sidebar-nav-item-store-notice',
      //       icon: 'StoreNoticeIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `${Routes.message.list}`,
      //       label: 'sidebar-nav-item-message',
      //       icon: 'ChatIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //   ],
      // },

      // feedback: {
      //   href: '',
      //   label: 'Feedback control',
      //   icon: 'SettingsIcon',
      //   childMenu: [
      //     {
      //       href: (shop: string) => `/${shop}${Routes.reviews.list}`,
      //       label: 'sidebar-nav-item-reviews',
      //       icon: 'ReviewIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //     {
      //       href: (shop: string) => `/${shop}${Routes.question.list}`,
      //       label: 'sidebar-nav-item-questions',
      //       icon: 'QuestionIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //   ],
      // },

      // user: {
      //   href: '',
      //   label: 'User control',
      //   icon: 'SettingsIcon',
      //   childMenu: [
      //     {
      //       href: (shop: string) => `/${shop}${Routes.staff.list}`,
      //       label: 'sidebar-nav-item-staffs',
      //       icon: 'UsersIcon',
      //       permissions: adminAndOwnerOnly,
      //     },
      //   ],
      // },

      promotional: {
        href: '',
        label: 'text-promotional-control',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.coupon.list}`,
            label: 'Coupons',
            icon: 'CouponsIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          {
            href: (shop: string) => `/${shop}${Routes.flashSale.list}`,
            label: 'text-flash-sale',
            icon: 'UsersIcon',
            childMenu: [
              {
                href: (shop: string) => `/${shop}${Routes.flashSale.list}`,
                label: 'text-available-flash-deals',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) =>
                  `/${shop}${Routes.myProductsInFlashSale}`,
                label: 'text-my-products-in-deals',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
              {
                href: (shop: string) =>
                  `/${shop}${Routes.vendorRequestForFlashSale.list}`,
                label: 'See all enrollment request',
                icon: 'ProductsIcon',
                permissions: adminOwnerAndStaffOnly,
              },
            ],
          },
        ],
      },

      layout: {
        href: '',
        label: 'text-page-management',
        icon: 'SettingsIcon',
        childMenu: [
          {
            href: (shop: string) => `/${shop}${Routes.faqs.list}`,
            label: 'text-faqs',
            icon: 'TypesIcon',
            permissions: adminOwnerAndStaffOnly,
          },
          // {
          //   href: (shop: string) => `/${shop}${Routes.termsAndCondition.list}`,
          //   label: 'Terms And Conditions',
          //   icon: 'TypesIcon',
          //   permissions: adminAndOwnerOnly,
          // },
        ],
      },
    },

    ownerDashboard: [
      {
        href: Routes.dashboard,
        label: 'sidebar-nav-item-dashboard',
        icon: 'DashboardIcon',
        permissions: ownerAndStaffOnly,
      },
      {
        href: Routes?.ownerDashboardMyShop,
        label: 'common:sidebar-nav-item-my-shops',
        icon: 'MyShopOwnerIcon',
        permissions: ownerAndStaffOnly,
      },
      {
        href: Routes?.ownerDashboardMessage,
        label: 'common:sidebar-nav-item-message',
        icon: 'ChatOwnerIcon',
        permissions: ownerAndStaffOnly,
      },
      {
        href: Routes?.ownerDashboardNotice,
        label: 'common:sidebar-nav-item-store-notice',
        icon: 'StoreNoticeOwnerIcon',
        permissions: ownerAndStaffOnly,
      },
      {
        href: Routes?.ownerDashboardShopTransferRequest,
        label: 'Shop Transfer Request',
        icon: 'MyShopIcon',
        permissions: adminAndOwnerOnly,
      },
    ],
  },
  product: {
    placeholder: '/product-placeholder.svg',
  },
  avatar: {
    placeholder: '/avatar-placeholder.svg',
  },
};

export const socialIcon = [
  {
    value: 'FacebookIcon',
    label: 'Facebook',
  },
  {
    value: 'InstagramIcon',
    label: 'Instagram',
  },
  {
    value: 'TwitterIcon',
    label: 'Twitter',
  },
  {
    value: 'YouTubeIcon',
    label: 'Youtube',
  },
];
