import * as yup from 'yup';

export const patientValidationSchema = yup.object().shape({
  name: yup.string().required('form:error-name-required'),
  age: yup
    .number()
    .typeError('form:error-age-must-number')
    .positive('form:error-age-must-positive')
    .required('form:error-age-required'),
  gender: yup.object().required('form:error-gender-required'),
});
