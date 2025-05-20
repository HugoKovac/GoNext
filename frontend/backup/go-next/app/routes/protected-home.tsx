import type { Route } from "../+types/home";

export function meta({}: Route.MetaArgs) {
  return [
    { title: "GoNext - Dashboard" },
    { name: "description", content: "Your GoNext Dashboard" },
  ];
}

export default function ProtectedHome() {
  return (
    <div className="container mx-auto px-4 py-8">
      <h1 className="text-4xl font-bold mb-8">Welcome Back!</h1>
      <div className="grid grid-cols-1 md:grid-cols-2 lg:grid-cols-3 gap-6">
        <div className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h2 className="card-title">Dashboard</h2>
            <p>This is your protected dashboard area.</p>
          </div>
        </div>
        <div className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h2 className="card-title">Profile</h2>
            <p>Manage your profile settings here.</p>
          </div>
        </div>
        <div className="card bg-base-100 shadow-xl">
          <div className="card-body">
            <h2 className="card-title">Settings</h2>
            <p>Configure your application settings.</p>
          </div>
        </div>
      </div>
    </div>
  );
} 