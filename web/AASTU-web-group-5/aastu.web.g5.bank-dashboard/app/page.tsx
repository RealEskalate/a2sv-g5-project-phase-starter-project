// app/page.tsx
'use client'

import React from 'react'
import Services from './Services/page'  // Ensure this path is correct
import Settings from './Settings/page 1/page'  // Ensure this path is correct

const Home = () => {
  return (
    <div> 
       {/*  <div>
        <Services/>
      </div>  */}
      <div> 
        <Settings/>
      </div> 
    </div>
  )
}

export default Home
