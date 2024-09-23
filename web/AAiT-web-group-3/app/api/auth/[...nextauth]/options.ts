import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from "next-auth/jwt";
import { Session } from "next-auth";

const BankLogin = CredentialsProvider({
  id: "bank-login",
  name: "Bank Login",
  credentials: {
    username: { label: "username", type: "text" },
    password: { label: "password", type: "password" },
  },
  async authorize(credentials) {
    if (!credentials) {
      throw new Error("Credentials are undefined");
    }

    const { username, password } = credentials;
    const userData = { userName: username, password: password };

    const result = await fetch(`${process.env.SERVER_URL}/auth/login`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    if (result.status === 200) {
      const resultData = await result.json();
      const { data } = resultData;
      return data;
    } else {
      return null;
    }
  },
});

const BankSignup = CredentialsProvider({
  id: "bank-signup",
  name: "Bank Signup",
  credentials: {
    name: { label: "name", type: "text" },
    email: { label: "email", type: "text" },
    username: { label: "username", type: "text" },
    password: { label: "password", type: "password" },
    confirmPassword: { label: "confirmPassword", type: "password" },
    role: { label: "role", type: "text" },
    // dateOfBirth: { label: "dateOfBirth", type: "text" },
    // permanentAddress: { label: "permanentAddress", type: "text" },
    // presentAddress: { label: "presentAddress", type: "text" },
    // postalCode: { label: "postalCode", type: "text" },
    city: { label: "city", type: "text" },
    // country: { label: "country", type: "text" },
    // profilePicture: { label: "profilePicture", type: "text" },
    // preference: {
    //   currency: "ETB",
    //   sentOrReceiveDigitalCurrency: true,
    //   receiveMerchantOrder: true,
    //   accountRecommendations: true,
    //   timeZone: "string",
    //   twoFactorAuthentication: true,
    // },
  },
  async authorize(credentials) {
    if (!credentials) {
      throw new Error("Credentials are undefined");
    }

    const { name, email, password, confirmPassword, role } = credentials;
    const userData = { email, password, name, confirmPassword, role };
    const result = await fetch(`${process.env.SERVER_URL}/signup`, {
      method: "POST",
      headers: {
        "Content-Type": "application/json",
      },
      body: JSON.stringify(userData),
    });

    if (result.status === 200) {
      const data = await result.json();
      if (data.message === "Successfully sent OTP") {
        return { email: email, id: "", name: null, role: "unverified" };
      }
    }
    return null;
  },
});

export const options = {
  providers: [BankLogin, BankSignup],
  callbacks: {
    async jwt({ token, user }: { token: JWT; user?: any }) {
      if (user) {
        token.accessToken = user.accessToken || "";
        token.refreshToken = user.refreshToken || "";
      }
      return token;
    },
    async session({ session, token }: { session: Session; token: JWT }) {
      if (session.user) {
        session.user.role = token.role;
        session.user.accessToken = token.accessToken;
        session.user.refreshToken = token.refreshToken;
      }
      return session;
    },
  },
  pages: {
    signIn: "/login",
    error: "/login",
  },
};
