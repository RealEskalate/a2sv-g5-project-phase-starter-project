"use client";
import { useSession } from "next-auth/react";
import { useDispatch, useSelector } from "react-redux";
import React, { useEffect } from "react";
import User from '../../type/user'
interface ExtendedUser {
	name?: string;
	email?: string;
	image?: string;
	accessToken?: string;
  }
export default function Login() {

  const { data: session, status } = useSession();
  const user = useSelector((state: { user: User }) => state.user);
  // console.log(session,'session')
  const users = session?.user as ExtendedUser;

  // console.log(users.accessToken,'accessToken');
  const dispatch = useDispatch();

  useEffect(() => {
    if (status === "authenticated" && users?.accessToken) {
      console.log('Dispatching USER_FETCH_REQUESTED');
      dispatch({
        type: "USER_FETCH_REQUESTED",
        payload: {
          userName: session.user.name,
          accessToken: users.accessToken,
        },
      });
    }
  }, [status, session, dispatch]);

  if (status === "loading") {
    return <div>Loading...</div>;
  }

  if (status === "unauthenticated") {
    return <div>Please log in</div>;
  }

  console.log('user', user);
  return <div> 
</div>

}