import React from 'react'
import EditProfile from '../components/EditProfile'
import TabSelector from '../components/TabSelector'
import Preferences from '../components/Preferences'
import Security from '../components/Security'

const page = () => {
  return (
    <div className='p-4'>
      <TabSelector tabs={["Edit Profile", "Preferences", "Security"]} contents={[<EditProfile />, <Preferences />, <Security/>]} />

    </div>
  )
}

export default page