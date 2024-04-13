import type { NextAuthConfig } from 'next-auth';
import { authorized } from './core/auth';
 
export const authConfig = {
  providers: [],
  pages: {
    signIn: '/login',
  },
  callbacks: {
    authorized: authorized,
  },
} satisfies NextAuthConfig;``