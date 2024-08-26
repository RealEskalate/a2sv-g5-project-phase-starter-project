"use client";
import React from "react";
import { useSession } from "next-auth/react";
import { useRouter } from "next/navigation";
import Landing from "./components/landing";

const HomePage: React.FC = () => {
  const session = useSession()
  const router = useRouter()

  if (session.data?.accessToken){
    router.push('/dashboard')
  }
  else{
    return(
      <Landing/>
    )
  }

};

export default HomePage;
