import axios from '../../../node_modules/axios'
export default {  
  debug: true,
  state: {
    name: '',
    icon_path: '',
    isUser: false,
  },
  set_info () {
    axios.get('https://localhost:8080/authUser', {
      // cookie
    }).then((res) => {
      this.state.name = res.data.Name
      this.state.icon_path = res.data.IconPath
      this.state.isUser = true
    })
    .catch((e) => { if(e.response) console.error('get error') })
  },
}
