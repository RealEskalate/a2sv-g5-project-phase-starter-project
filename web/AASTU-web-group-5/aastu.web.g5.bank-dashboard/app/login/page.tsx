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
  const users = session?.user as ExtendedUser;

  const dispatch = useDispatch();
  
  useEffect(() => {
    if (status === "authenticated" && users?.accessToken) {
      console.log('Dispatching USER_FETCH_REQUESTED');
      dispatch({
        type: "USER_FETCH_REQUESTED",
        payload: {
          username: users.name,
          token: users.accessToken,
        },
      });
    }
  }, [status, session, dispatch, users]);
  useEffect(() => {
    console.log('Updated user state:', user); // Log the updated user state
  }, [user]);

  // return <div> 
// </div>
return (<></>)
}