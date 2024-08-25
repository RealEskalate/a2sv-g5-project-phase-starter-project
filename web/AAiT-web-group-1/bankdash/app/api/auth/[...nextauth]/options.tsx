import CredentialsProvider from "next-auth/providers/credentials";

export const options = {
  providers: [
    CredentialsProvider({
        name: "Credentials",
        credentials: {
        userName: { label: 'User Name', type: 'text' },
        password: { label: 'Password', type: 'password' },
      },
      
      async authorize(credentials:any) {
        console.log("first")
        const response = await fetch('https://bank-aait-web-group-1.onrender.com/auth/login', {
          method: 'POST',
          headers: { 'Content-Type': 'application/json' },
          body: JSON.stringify({
            userName: credentials.userName,
            password: credentials.password,
          }),
        });

        const data = await response.json();

        if (data.success && data.data) {
          return {
            id: data.data.id,
            name: data.data.name,
            email: data.data.email,
            accessToken: data.data.accessToken,
            refreshToken: data.data.refreshToken,
          };
        }

        return null;
      },
    }),
  ],
  callbacks: {
    async jwt({ token, user }: { token: any, user: any }) {
      if (user) {
        token.accessToken = user.accessToken;
        token.profileStatus = user.profileStatus;
      }
      return token;
    },
    async session({ session, token }: { session: any, token: any }) {
      session.user.accessToken = token.accessToken;
      session.user.profileStatus = token.profileStatus;
      return session;
    },
  },
  pages: {
    signIn: '/auth/login',  
  },
};