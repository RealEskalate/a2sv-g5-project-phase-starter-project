import CredentialsProvider from "next-auth/providers/credentials";

import Provider from "next-auth";

export const Options = {
  providers: [
    CredentialsProvider({
      name: "credentials",
      credentials: {
        userName: { label: "UserName", placehlder: "username", type: "text" },
        password: {
          label: "Password",
          placeholder: "enter your password",
          type: "password",
        },
      },
      async authorize(credentials) {
        console.log(credentials);
        // Add your authorization logic here
        return null; // or return a User object based on your application logic
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
