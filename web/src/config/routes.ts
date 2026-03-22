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
  verifyLicense: '/verify-license',
  user: {
    ...routesFactory('/users'),
  },
  shop: {
    ...routesFactory('/shops'),
  },
  order: {
    ...routesFactory('/orders'),
  },
  withdraw: {
    ...routesFactory('/withdraws'),
  },
  newShops: '/new-shops',
  draftProducts: '/products/draft',
  outOfStockOrLowProducts: '/products/product-stock',
  productInventory: '/products/inventory',
  transaction: '/orders/transaction',
  termsAndCondition: {
    ...routesFactory('/terms-and-conditions'),
  },
  adminList: '/users/admins',
  vendorList: '/users/vendors',
  pendingVendorList: '/users/vendors/pending',
  customerList: '/users/customer',
  myStaffs: '/users/my-staffs',
  vendorStaffs: '/users/vendor-staffs',
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
  flashSale: {
    ...routesFactory('/flash-sale'),
  },
  ownerDashboardNotice: '/notice',
  ownerDashboardMessage: '/owner-message',
  ownerDashboardMyShop: '/my-shop',
  myProductsInFlashSale: '/flash-sale/my-products',
  ownerDashboardNotifyLogs: '/notify-logs',
  inventory: {
    editWithoutLang: (slug: string, shop?: string) => {
      return shop ? `/${shop}/products/${slug}/edit` : `/products/${slug}/edit`;
    },
    edit: (slug: string, language: string, shop?: string) => {
      return shop
        ? `/${language}/${shop}/products/${slug}/edit`
        : `/${language}/products/${slug}/edit`;
    },
    translate: (slug: string, language: string, shop?: string) => {
      return shop
        ? `/${language}/${shop}/products/${slug}/translate`
        : `/${language}/products/${slug}/translate`;
    },
  },
  visitStore: (slug: string) => `${process.env.NEXT_PUBLIC_SHOP_URL}/${slug}`,
  vendorRequestForFlashSale: {
    ...routesFactory('/flash-sale/vendor-request'),
  },
  becomeSeller: '/become-seller',
  ownershipTransferRequest: {
    ...routesFactory('/shop-transfer'),
  },
  ownerDashboardShopTransferRequest: '/shop-transfer/vendor',
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
