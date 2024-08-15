import React, { Suspense } from 'react'
import Form from './component/Form'

const SigninPage = () => {
  return (
    <div>
      <Suspense>
       <Form/>
      </Suspense>
    </div>
  )
}

export default SigninPage



