'use client';

import { Geist, Geist_Mono } from "next/font/google";
import "./globals.css";
import Navbar from "./components/navbar";
import { useState, useEffect } from "react";

const geistSans = Geist({
  variable: "--font-geist-sans",
  subsets: ["latin"],
});

const geistMono = Geist_Mono({
  variable: "--font-geist-mono",
  subsets: ["latin"],
});

export default function RootLayout({ children }) {
  const [user, setUser] = useState(null);

  useEffect(() => {
    const token = document.cookie.split('; ').find(row => row.startsWith('token=')).split('=')[1];
    if (token) {
      setUser(token);
    }
  }, []);
  return (
    <html lang="en">
      <body
        className={`${geistSans.variable} ${geistMono.variable} antialiased`}
      >
        <Navbar user={user} />
        {children}
      </body>
    </html>
  );
}
