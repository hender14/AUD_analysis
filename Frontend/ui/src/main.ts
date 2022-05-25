import { createApp } from 'vue'
import App from './App.vue'
import router from './router'
import store from './store'
import axios from 'axios'
import { initializeApp } from "firebase/app";
import { getStorage, ref, uploadBytes } from "firebase/storage";

export const userURL = process.env.VUE_APP_USERURL
export const analysisURL = process.env.VUE_APP_ANALYSISURL

axios.defaults.withCredentials = true

// Your web app's Firebase configuration
const firebaseConfig = {
  apiKey: process.env.VUE_APP_APIKEY,
  authDomain: process.env.VUE_APP_AUTHDOMAIN,
  projectId: process.env.VUE_APP_PROJECTID,
  storageBucket: process.env.VUE_APP_STORAGEBUCKET,
  messagingSenderId: process.env.VUE_APP_MESSAGINGSENDERID,
  appId: process.env.VUE_APP_APPID
};

export async function jsondec(data: Promise<ArrayBuffer>) {
  const check = new TextDecoder().decode(await data)
  const json = JSON.parse(check.toString());
  return json
}

export async function uploadfile(file: any, username: string, filename: string) {
  const bucketName = process.env.VUE_APP_BUCKETFORDER
  const time = new Date().getTime().toString();
  const fileName = time+'-'+filename
  const pathName = 'gs://'+ firebaseConfig.storageBucket + '/' + bucketName+'/'+ username+'/'+ fileName

  // Initialize Firebase
  const firebaseApp = initializeApp(firebaseConfig);
  const storage = getStorage(firebaseApp);
  const gsReference = ref(storage, pathName);
  await uploadBytes(gsReference, file)

  return fileName
}

createApp(App).use(store).use(router).mount('#app')