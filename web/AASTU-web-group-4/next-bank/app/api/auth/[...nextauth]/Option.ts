import CredentialsProvider from "next-auth/providers/credentials";
import { store } from "../../../../lib/redux/store";
import { authApi } from "../../../../lib/redux/features/authapi";
import Provider from "next-auth";

export const Options = {
  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        username: { label: "Username", placehlder: "username", type: "text" },
        password: {
          label: "Password",
          placeholder: "enter your password",
          type: "password",
        },
      },
      async authorize(credentials) {
        const result = await store.dispatch(
          authApi.endpoints.login.initiate(credentials)
        );

        if (result.data) {
          return result.data;
        } else {
          return null;
        }
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }: { token: any; user: any }) {
      if (user) {
        token.role = user.role;
        token.accessToken = user.data.accessToken;
        token.refreshToken = user.refreshToken;
        token.name = user.data.name;
        token.email = user.data.email;
        token.username = user.data.username;
        token.postalCode = user.data.postalCode;
        token.profilePicture = user.data.profilePicture;
        token.currency = user.data.preference.currency;
      }

      return token;
    },
    async session({ session, token }: { session: any; token: any }) {
      if (session?.user) {
        session.user.role = token.role;
        session.accessToken = token.accessToken;
        session.refreshToken = token.refreshToken;
        session.user.name = token.name;
        session.user.email = token.email;
        session.user.username = token.username;
        session.user.postalCode = token.postalCode;
        session.user.profilePicture = token.profilePicture;
        session.user.currency = token.currency;
      }
      return session;
    },
    // console.log(session, "when sesssion called");
    // return session;
  },
  pages: {
    signIn: "/signin",
    // signUp:'/signup',
  },
};
