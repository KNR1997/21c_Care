import * as yup from 'yup';

export const patientValidationSchema = yup.object().shape({
  name: yup.string().required('form:error-name-required'),
  age: yup.string().required('form:error-age-required'),
  gender: yup.string().required('form:error-gender-required'),
});
