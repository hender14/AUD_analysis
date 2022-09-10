<template>
  <!-- <h1>{{ message }}</h1> -->
  <div v-if="notify.cls" :class="`alert alert-${notify.cls}`" role="alert">
  {{ notify.message }}
  </div>
  <header class="border-bottom pb-2 mb-3 d- flex align-items-center">
    <h1 class="fs-5 m-0">音声解析</h1>
  </header>
  <form @submit.prevent="upload">
    <div>
      <label class="pb-2 mb-3 d- flex align-items-center" for="file">アップロードするファイルを選択してください</label><br>
      <input class="pb-2 mb-3 d- flex align-items-center" type="file" id="file" name="file" required @change="validate_uploads"><br>
      <input type='submit' name='upload_btn' value='upload'>
    </div>
  </form>
</template>


<script>
import { reactive, ref, onMounted } from 'vue'
import axios from 'axios'
import { useRouter } from 'vue-router'
import { useStore } from 'vuex'
import { uploadfile, userURL, analysisURL } from '@/main'

export default {
  name: "Result",
  setup() {
    const message = ref('You are not logged in!')
    let file,
    target,
    name,
    size,
    type,
    errors = ref(''),
    user
    const notify = reactive({
      cls: '',
      message: ''
    })

    const validate_uploads = async (event) => {
      file = event.target.files[0],
      name = file.name,
      size = file.size,
      type = file.type,
      errors = ref('')
      // let date = new Date()
      // var time = date.getFullYear() + (date.getMonth()+1) + date.getDate() + date.getHours() + date.getMinutes() + date.getSeconds()

        var reader = new FileReader();
        reader.onload = function(event){
          target = event.target
        }
        reader.readAsArrayBuffer(file);
        // reader.readAsDataURL(file);
        // reader.readAsText(file, "Shift_JIS");

      //上限サイズは30MB
      if (size > 30000000) {
        errors.value += 'ファイルの上限サイズ30MBを超えています\n'
      }

      //拡張子は .jpg .gif .png . pdf のみ許可
      // if (type != 'text/plain' && type != 'audio/*' ) {
      //   errors.value += '.txt、audio/*のいずれかのファイルのみ許可されています\n'
      // }

      if (errors.value) {
        //errorsが存在する場合は内容をalert
        alert(errors.value)
        //valueを空にしてリセットする
        event.currentTarget.value = ''
      }
      message.value = ref('file validate')
    }

    const store = useStore()
    const router = useRouter()

    const upload = async () => {
      try {
        // parameter の用意
        const time = new Date().getTime().toString();
        const filename = time+'-'+name
        const params = new FormData();
        params.append("file", file);
        params.append("username", user);
        params.append("filename", filename);
        axios.post( analysisURL + 'analysis', params, {
          headers: {
            // multipartで送信
            'content-type': 'multipart/form-data',
          },
        }).then(res => {
          console.log(res)
        // alert("「" + res['name'] + "」登録完了");
        // this.$router.push({path: '/articles/list'});
        })
        notify.cls = 'success'
        notify.message = 'file upload completed!'
      } catch (e) {
        // if (axios.isAxiosError(e) && e.response && e.response.status === 400) {
        if (axios.isAxiosError(e) && e.response.status === 400) {
          notify.cls = 'danger'
          notify.message = e.response.data["message"]
          // message.value = e.response.data["message"]
        }
      }
    }

    onMounted(async () => {
      try {
        const { data } = await axios.get( userURL+ 'user')
        user = data.ID
        // actionsに設定したパラメータ名を設定
        await store.dispatch('setAuth', true)
      } catch(e) {
        await store.dispatch('setAuth', false)
        router.push('/login')
      }
    })


    return {
      message,
      validate_uploads,
      upload,
      notify,
    }
  }

}
</script>