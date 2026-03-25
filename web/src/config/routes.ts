export const Routes = {
  dashboard: '/',
  login: '/login',
  logout: '/logout',
  register: '/register',
  forgotPassword: '/forgot-password',
  resetPassword: '/reset-password',
  profile: '/profile',
  settings: '/settings',
  seoSettings: '/settings/seo',
  profileUpdate: '/profile-update',
  verifyEmail: '/verify-email',
  user: {
    ...routesFactory('/users'),
  },
  adminList: '/users/admins',
  patient: {
    ...routesFactory('/patients'),
  },
  drug: {
    ...routesFactory('/drugs'),
  },
  labTest: {
    ...routesFactory('/labTests'),
  },
  visit: {
    ...routesFactory('/visits'),
  },
  billing: {
    ...routesFactory('/billings'),
  },
  ownerDashboardNotice: '/notice',
  ownerDashboardMessage: '/owner-message',
  ownerDashboardMyShop: '/my-shop',
  myProductsInFlashSale: '/flash-sale/my-products',
  ownerDashboardNotifyLogs: '/notify-logs',
  visitStore: (slug: string) => `${process.env.NEXT_PUBLIC_SHOP_URL}/${slug}`,
};

function routesFactory(endpoint: string) {
  return {
    list: `${endpoint}`,
    create: `${endpoint}/create`,
    editWithoutLang: (slug: string, shop?: string) => {
      return shop
        ? `/${shop}${endpoint}/${slug}/edit`
        : `${endpoint}/${slug}/edit`;
    },
    edit: (slug: string, language: string, shop?: string) => {
      return shop
        ? `/${language}/${shop}${endpoint}/${slug}/edit`
        : `${language}${endpoint}/${slug}/edit`;
    },
    translate: (slug: string, language: string, shop?: string) => {
      return shop
        ? `/${language}/${shop}${endpoint}/${slug}/translate`
        : `${language}${endpoint}/${slug}/translate`;
    },
    details: (slug: string) => `${endpoint}/${slug}`,
    editByIdWithoutLang: (id: string, shop?: string) => {
      return shop ? `/${shop}${endpoint}/${id}/edit` : `${endpoint}/${id}/edit`;
    },
  };
}
