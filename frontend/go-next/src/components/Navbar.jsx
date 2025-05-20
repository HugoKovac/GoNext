import { useLocation } from "react-router";
import { useAuth } from "../context/AuthContext";

function Navbar() {
  const { user, isLoading, logout } = useAuth();

  return (
    <div className="navbar bg-base-100 shadow-sm">
      <div className="navbar-start">
        <div className="dropdown">
          <div tabIndex={0} role="button" className="btn btn-ghost btn-circle">
            <svg
              xmlns="http://www.w3.org/2000/svg"
              className="h-5 w-5"
              fill="none"
              viewBox="0 0 24 24"
              stroke="currentColor"
            >
              {" "}
              <path
                strokeLinecap="round"
                strokeLinejoin="round"
                strokeWidth="2"
                d="M4 6h16M4 12h16M4 18h7"
              />{" "}
            </svg>
          </div>
          <ul
            tabIndex={0}
            className="menu menu-sm dropdown-content bg-base-100 rounded-box z-1 mt-3 w-52 p-2 shadow"
          >
            <li>
              <a>Homepage</a>
            </li>
            <li>
              <a>Portfolio</a>
            </li>
            <li>
              <a>About</a>
            </li>
            {!user && (
              <li className="md:hidden">
                <a href="/signin">Sign In</a>
              </li>
            )}
            {!user && (
              <li className="md:hidden">
                <a href="/signup">Sign Up</a>
              </li>
            )}
          </ul>
        </div>
      </div>
      <div className="navbar-center">
        <a className="btn btn-ghost text-xl" href="/">
          GoNext
        </a>
      </div>
      <div className="navbar-end">
        <div className="hidden md:flex">
          {!user && (
            <a className="btn btn-primary mx-2" href="/signin">
              Sign In
            </a>
          )}
          {!user && (
            <a className="btn mx-2" href="/signup">
              Sign Up
            </a>
          )}
          {user && (
            <button className="btn mx-2" onClick={logout}>
              Sign Out
            </button>
          )}
        </div>
      </div>
    </div>
  );
} 

export default Navbar;