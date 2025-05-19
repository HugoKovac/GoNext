/** @type {import('next').NextConfig} */
const nextConfig = {
  async rewrites() {
    return [
      {
        source: '/api/:path*',
        destination: 'http://backend_dev:8080/api/:path*',
      },
    ];
  },
};

export default nextConfig;
