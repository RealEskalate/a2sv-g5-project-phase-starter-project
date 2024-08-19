import { getServerSession } from "next-auth";
import { Button } from "@/components/ui/button";
import Link from "next/link";
import { authOptions } from "./api/auth/[...nextauth]/options";
import { redirect } from "next/navigation";

const Home = async () => {
  const session = await getServerSession(authOptions);
  if (!session) {
    redirect("/auth/sign-in");
  }
  return <></>;
};

export default Home;
