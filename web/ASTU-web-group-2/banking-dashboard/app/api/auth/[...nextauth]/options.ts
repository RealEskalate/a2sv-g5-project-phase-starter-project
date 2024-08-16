import CredentialsProvider from "next-auth/providers/credentials";
import fetchData from "../../fetch";

export const options = {
    providers: [
        CredentialsProvider({
            type: "credentials",
            credentials: {
                userName: { label: "userName", type: "text", placeholder: "jsmith" },
                password: { label: "Password", type: "password" },
            },
            async authorize(credentials) {
                const { userName, password } = credentials as {
                    userName: string;
                    password: string;
                };

                try {
                    const response = await fetchData({ userName, password });
                    console.log("Response:", response);
                    if (response && response.id) {
                        return {
                            id: response.id,
                            accessToken: response.accessToken,
                            refreshToken: response.refreshToken,
                        };
                    }
                } catch (err) {
                    console.error('Failed to login:', err);
                }

                // Return null if authentication fails
                return null;
            },
        }),
    ],
    pages: {
        signIn: '/login',
    },
    callbacks: {
        async jwt({ token, user }: { token: any, user: any }) {
            if (user) {
                token.accessToken = user.accessToken; // Ensure accessToken is stored
                token.refreshToken = user.refreshToken; // Store refreshToken
            }
            return token;
        },

        async session({ session, token }: { session: any, token: any }) {
            if (session?.user) {
                session.user.accessToken = token.accessToken;
                session.user.refreshToken = token.refreshToken; 
            }
            return session;
        },

        // async redirect({ url, baseUrl }: { url: string, baseUrl: string }) {
        //     return baseUrl; 
        // },
    },
};
