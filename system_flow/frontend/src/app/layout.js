import { Inter } from "next/font/google";
import "./globals.css";
import { AppWrapper } from "@/context/appContext";

const inter = Inter({ subsets: ["latin"] });

export const metadata = {
  title: "Todolist app",
  description: "Case study for sprint asia",
};

export default function RootLayout({ children }) {
  return (
    <html lang="en">
      <body className={inter.className}>
        <AppWrapper>{children}</AppWrapper>
      </body>
    </html>
  );
}
