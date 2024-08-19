import { withAuth } from 'next-auth/middleware';

export default withAuth({
  pages: {
    signIn: '/component/signin', 
  },
})


export const config = {
    matcher: ['/:path*'],
};