import * as yup from 'yup';

export const drugValidationSchema = yup.object().shape({
  name: yup.string().required('form:error-name-required'),
  default_price: yup.string().required('form:error-price-required'),
});
