import NextAuth from 'next-auth';
import CredentialsProvider from 'next-auth/providers/credentials';
import { bankApi } from '@/lib/redux/api/bankApi';
import { store } from '@/lib/redux/store';
import { User } from '@/types/User';
import { UserResponse } from '@/lib/redux/types/userResponse';
import { UserSignup } from '@/types/UserSignUp';




const handler = NextAuth({
  // session: {
  //   strategy: 'jwt'
  // },
  providers: [
    CredentialsProvider({
      id: 'sign-up',
      credentials: {
        name : {type: 'text'},
        email : {type: 'email'},
        dateOfBirth : {type: 'Date'},
        permanentAddress : {type: 'text'},
        postalCode : {type: 'text'},
        username : {type: 'text'},
        password: {type: "password"},
        presentAddress : {type: 'text'},
        city : {type: 'text'},
        country : {type: 'text'},

      },
      authorize: async (credentials) => {
        const userData = { 
          email: credentials?.email,
          name: credentials?.name,
          password: credentials?.password,
          username: credentials?.username,
          presentAddress: credentials?.presentAddress,
          dateOfBirth: credentials?.dateOfBirth,
          permanentAddress: credentials?.permanentAddress,
          postalCode: credentials?.postalCode,
          country: credentials?.country,
          city: credentials?.city,
          preference: {
          currency: "birr",
          sentOrReceiveDigitalCurrency: true,
          receiveMerchantOrder: true,
          accountRecommendations: true,
          timeZone: "EAT",
          twoFactorAuthentication: true,

        }}
        console.log("signup userData",credentials, userData)

        const response = await store.dispatch(bankApi.endpoints.signup.initiate(userData as UserSignup))
      
        console.log("response authorize", response)
        const res = response.data
        if (res?.success) {
          console.log("success authorize", res, res.data)
          return {data: res.data.data, access_token: res.data.access_token, refresh_token: res.data.refresh_token, id: res.data.data.id };
          
        } else {
          console.log("failure authorize null", res)

          return null
        }
      },
    }),
    CredentialsProvider({
      id: 'sign-in',
      credentials: {
        userName: {label: 'Username', type:'text'},
        password: { label: 'Password', type: 'password' },

      },
      authorize: async (credentials) => {
        const loginData = {
            userName: credentials?.userName,
            password: credentials?.password
          }

        const response = await store.dispatch(bankApi.endpoints.signin.initiate(loginData as SigninCredential))
        const res = response.data

        if (res?.success && res.data) {
          return {access_token: res.data.access_token, refresh_token: res.data.refresh_token };

        } else {
          return null
        }
      },
    }),
  ],
  pages:{
    signIn: '/auth/signin',
    signOut: '/auth/signout',
    error: '/auth/error',
    verifyRequest: '/auth/verify-request',
    newUser: '/auth/signup'
  },
  callbacks: {
    async jwt({ token, user, account }) {
      console.log("jwt user", user, "token", token)
      if (user) {
        token.access_token = user.access_token;
        token.refresh_token = user.refresh_token;
        if(account?.provider === "sign-up")
          token.user = user.data;

      }
      console.log("jwt token", token)
      return token;
    },
    async session({ session, token, user }) {
      console.log("jwt user", user, "token", token, "serssion", session)

      if (token) {
        session.user = token?.user as User;
        session.access_token = token.access_token as string;
        session.refresh_token = token.refresh_token as string;

        
      }
      console.log("jwt session", session)

      return session;
    }
  }
});

export { handler as GET, handler as POST };
