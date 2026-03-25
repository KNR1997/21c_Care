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
