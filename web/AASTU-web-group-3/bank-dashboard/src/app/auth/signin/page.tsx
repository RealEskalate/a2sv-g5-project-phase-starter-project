import React, { Suspense } from 'react'
import Form from './component/Form'

const SigninPage = () => {
  return (
    <div className= ' h-screen flex items-center justify-center'>
      <Suspense>
       <Form/>
      </Suspense>
    </div>
  )
}

export default SigninPage



