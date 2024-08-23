"use client";
import React, { useEffect } from "react";
import Head from "next/head";
import Navbar from "../components/navbar/Navbar";
import Sidebar from "../components/sidebar/Sidebar";
import { Inter } from "next/font/google";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation"; // Correct import
import { useGetCurrentUserQuery } from "@/lib/service/UserService";
import { useDispatch, useSelector } from "react-redux";
import { RootState, AppDispatch } from "@/lib/store";
import { setUser } from "@/lib/features/userSlice/userSlice";
import { isTokenExpired } from "@/utils/authUtils";
import { useRefreshAccessTokenMutation } from "@/lib/service/authentication";
import LayoutForTest from "../components/layout/LayoutForTest";

const inter = Inter({ subsets: ["latin"] });

const Layout = ({
  children,
  title = "My Next.js App",
}: {
  children: React.ReactNode;
  title?: string;
}) => {
  const { data: session, status } = useSession();
  const router = useRouter();
  const dispatch = useDispatch<AppDispatch>();
  const user = useSelector((state: RootState) => state.user.user);

  const [refreshAccessToken, { isLoading: isRefreshing }] =
    useRefreshAccessTokenMutation();

  const { data: userData, isLoading } = useGetCurrentUserQuery(
    session?.user?.accessToken ?? "",
    {
      skip: !session?.user?.accessToken,
    }
  );

  useEffect(() => {
    const fetchUserData = async () => {
      try {
        const refreshToken = session?.user?.refreshToken;
        if (!refreshToken) {
          throw new Error("No refresh token available");
        }

        const refreshData = await refreshAccessToken(refreshToken);
        const newAccessToken = refreshData.data.data;

        if (newAccessToken) {
          session.user.accessToken = newAccessToken;
        } else {
          throw new Error("Failed to refresh access token");
        }
      } catch (error) {
        console.error("Failed to refresh access token:", error);
        router.push("/login");
      }
    };
    if (session) {
      console.log(isTokenExpired(session.user.accessToken));
    }
    console.log("session && isTokenExpired(session.user.accessToken)",session && isTokenExpired(session.user.accessToken))
    // if (session && isTokenExpired(session.user.accessToken)) {
      // fetchUserData();
    // }
    // if (session?.user?.accessToken) {
      // console.log("session", session, isTokenExpired(session.user.accessToken));
    // }
    if (status == "unauthenticated" && !session) {
      router.push("/login");
    }
  }, [session, router, refreshAccessToken]);

  useEffect(() => {
    if (userData?.data) {
      dispatch(setUser(userData.data));
    }
  }, [userData, dispatch]);

  return (
    <>
      <Head>
        <title>{title}</title>
        <meta name="viewport" content="width=device-width, initial-scale=1" />
      </Head>
      <div className={`${inter.className} flex flex-col min-h-screen`}>
        <div className="flex flex-1">
          <LayoutForTest/>
          <main className="max-md:pt-[100px] flex-1 p-4 mt-[60px] lg:ml-[240px] md:ml-[240px] ml-0 bg-[#F5F7FA]">
            {children}
          </main>
        </div>
      </div>
    </>
  );
};

export default Layout;
