import React from 'react'
import EditProfile from '../components/EditProfile'
import TabSelector from '../components/TabSelector'

const page = () => {
  return (
    <div className='p-4'>
      <TabSelector tabs={["Edit Profile"]} contents={[<EditProfile />]} />

    </div>
  )
}

export default page