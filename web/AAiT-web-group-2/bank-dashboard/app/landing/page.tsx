"use client"

import { useSession } from "next-auth/react"

const page = () => {
    const session = useSession()
  return (
    <div>
        <p>Hi {session.data?.user?.name} </p>
        {
            session.data?.access_token
            


        }

    </div>
  )
}

export default page