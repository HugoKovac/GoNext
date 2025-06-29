import './App.css'

function App() {
  return (
    <div className="stack-info max-w-xl mx-auto p-8">
      <h1 className="text-4xl font-bold mb-4">🚀 Welcome to GoNext!</h1>
      <p className="text-lg mb-6">
        Imagine launching your next big idea in days, not weeks. <b>GoNext</b> is the all-in-one starter kit that empowers you to build, test, and deploy modern web applications with confidence and speed.
      </p>
      <ul className="text-base mb-6 list-disc list-inside">
        <li><b>Lightning-fast Backend:</b> Built with Golang & Ent ORM for performance and reliability.</li>
        <li><b>Modern Frontend:</b> ReactJS (Vite) for instant feedback and a delightful developer experience.</li>
        <li><b>Beautiful UI:</b> Styled with Tailwind CSS and DaisyUI for rapid, consistent, and customizable design.</li>
        <li><b>Rock-solid Database:</b> PostgreSQL, trusted by startups and enterprises alike.</li>
        <li><b>Seamless Dev Experience:</b> Spin up the entire stack with Docker Compose in seconds.</li>
        <li><b>Production Ready:</b> Effortless AWS deployment (EC2, RDS, S3, CloudFront, ALB, Terraform).</li>
        <li><b>Secure by Default:</b> JWT authentication, strong password validation, and best practices baked in.</li>
      </ul>
      <p className="text-base mb-6">
        <b>This very website is powered by GoNext.</b> Want to see it in action? <span className="text-primary">Sign up and log in</span> to explore the backend features, test authentication, and experience a real-world, production-grade stack.
      </p>
      <div className="bg-base-200 rounded-lg p-4 mb-6 border border-base-300">
        <b>Why choose GoNext?</b>
        <ul className="mt-2">
          <li>⏱️ Save weeks of setup and configuration</li>
          <li>🔒 Security and scalability from day one</li>
          <li>🛠️ Built for both rapid prototyping and serious production</li>
          <li>🌍 Open-source and ready for your next project</li>
        </ul>
      </div>
      <p className="text-base text-slate-600">
        Ready to build something amazing? <b>Try GoNext now!</b>
      </p>
    </div>
  )
}

export default App
