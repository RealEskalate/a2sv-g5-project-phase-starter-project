import { z } from 'zod';

export const formSchema = z.object({
  name: z.string().min(2, {
    message: "Name must be at least 2 characters.",
  }),
  email: z.string().email({
    message: "A valid email must be provided.",
  }),
  dateOfBirth: z.string().optional(),
  permanentAddress: z.string().min(5, {
    message: "Permanent Address must be at least 5 characters.",
  }),
  postalCode: z.string().min(4, {
    message: "Postal Code must be at least 4 characters.",
  }),
  username: z.string().min(2, {
    message: "Username must be at least 2 characters.",
  }),
  password: z.string().min(6, {
    message: "Password must be at least 6 characters.",
  }),
  presentAddress: z.string().min(5, {
    message: "Present Address must be at least 5 characters.",
  }),
  city: z.string().min(2, {
    message: "City must be at least 2 characters.",
  }),
  country: z.string().min(2, {
    message: "Country must be at least 2 characters.",
  }),
});


export const signUpSchema = z.object({ 
  name: z.string().min(2).max(50),
  email: z.string().email(),
  dateOfBirth: z.string(),
  permanentAddress: z.string(),
  postalCode: z.string(),
  username: z.string().min(2).max(50),
  password: z.string().min(8),
  presentAddress: z.string(),
  city: z.string(),
  country: z.string(),
  profilePicture: z.string(),
  preference: z.object({
    currency: z.string(),
    sentOrReceiveDigitalCurrency: z.boolean(),
    receiveMerchantOrder: z.boolean(),
    accountRecommendations: z.boolean(),
    timeZone: z.string(),
    twoFactorAuthentication: z.boolean(),
  }),
})


export const signInSchema = z.object({
  username: z.string().min(2).max(50),
  password: z.string().min(8),
});
