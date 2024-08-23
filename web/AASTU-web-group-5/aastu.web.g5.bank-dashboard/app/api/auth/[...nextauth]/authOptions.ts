import NextAuth, { User as NextAuthUser } from "next-auth";
import CredentialsProvider from "next-auth/providers/credentials";
import { JWT } from "next-auth/jwt"; // Import the JWT type
import creditCardColor from "@/app/CreditCards/cardMockData";
console.log('authOptions')
export const authOptions = {
	providers: [
		CredentialsProvider({
			name: "Credentials",
			credentials: {
				userName: { label: "Username", type: "text", placeholder: "" },
				password: { label: "Password", type: "password" },
			},
			async authorize(credentials) {
				console.log(creditCardColor,'credentails')
				const res = await fetch(
					"https://bank-dashboard-1tst.onrender.com/auth/login",
					{
						method: "POST",
						headers: {
							"Content-Type": "application/json",
						},
						body: JSON.stringify({
							userName: credentials?.userName,
							password: credentials?.password,
						}),
					}
				);

				const user = await res.json();

				if (user) {
					return {
						id: "f",
						username: credentials?.userName,
						accessToken: user.data.access_token,
						refreshToken: user.data.refresh_token,
					};
				} else {
					return null;
				}
			},
		}),
	],
	secret: process.env.NEXTAUTH_SECRET,
	callbacks: {
		async session({ session, token }: { session: any; token: JWT }) {
			if (token) {
				session.user.name = token.name;
				session.user.accessToken = token.accessToken;
				session.user.refreshToken = token.refreshToken;
			}
			console.log("sesssssssssss: ", session);
			return session;
		},
		async jwt({ token, user }: { token: JWT; user }) {
			if (user) {
				token.name = user.username;
				token.accessToken = user.accessToken;
				token.refreshToken = user.refreshToken;
			}
			return token;
		},
	},
};
