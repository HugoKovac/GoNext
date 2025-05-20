import type { Route } from "../+types/home";
import { Welcome } from "../welcome/welcome";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "GoNext - Welcome" },
    { name: "description", content: "Welcome to GoNext!" },
  ];
}

export default function PublicHome() {
  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-4xl font-bold mb-8">Welcome to GoNext</h1>
      <p className="text-lg mb-4">Please sign in to access the full features.</p>
      <Welcome />
    </div>
  );
} 