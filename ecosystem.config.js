module.exports = {
  apps : [{
    name   : "resize",
    script : "make build && ./resize",
    // watch: true,
    // kill_timeout : 1000,
    exec_interpreter: "none",
    exec_mode : "fork_mode",
  //  env_development: {
  //     NODE_ENV: "development"
  //  }
  "ignore_watch": ["resize"],
  watch: ["src"]
  }]
}
// pm2 start ecosystem.config.js 
// pm2 delete all
// need pm2 save to sync