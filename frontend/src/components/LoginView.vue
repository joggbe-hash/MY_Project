<template>
  <div class="auth-view-container">
    <!-- View Switcher Tabs on top for easy testing -->
    <div class="auth-nav-switch">
      <button 
        class="auth-mode-btn" 
        :class="{ active: authMode === 'login' }"
        @click="authMode = 'login'"
      >
        登入帳號
      </button>
      <button 
        class="auth-mode-btn" 
        :class="{ active: authMode === 'register' }"
        @click="authMode = 'register'"
      >
        建立帳號
      </button>
      <button 
        class="auth-mode-btn" 
        :class="{ active: authMode === 'forgot' }"
        @click="authMode = 'forgot'"
      >
        忘記密碼
      </button>
    </div>

    <!-- Alert / Toast -->
    <div v-if="alertMessage" class="auth-alert-box">
      {{ alertMessage }}
    </div>

    <!-- SCREEN 12: 登入帳號 -->
    <div v-if="authMode === 'login'" class="auth-screen-card">
      <div class="brand-header">
        <h1 class="brand-title">我也要</h1>
        <p class="brand-sub">把靈感變成行動，留下自己的成果。</p>
      </div>

      <div class="auth-fields-block">
        <h2 class="form-title">登入帳號</h2>

        <div class="form-field-group">
          <label class="field-label">電子郵件</label>
          <div class="field-input-box">
            <input 
              v-model="loginEmail" 
              type="email" 
              placeholder="輸入電子郵件" 
              class="auth-input"
            />
          </div>
        </div>

        <div class="form-field-group">
          <label class="field-label">密碼</label>
          <div class="field-input-box">
            <input 
              v-model="loginPassword" 
              type="password" 
              placeholder="輸入密碼" 
              class="auth-input"
              @keyup.enter="handleLogin"
            />
          </div>
        </div>

        <div class="forgot-link-row">
          <button type="button" class="text-link-btn" @click="authMode = 'forgot'">
            忘記密碼？
          </button>
        </div>

        <button type="button" class="btn-primary-auth" @click="handleLogin">
          登入
        </button>

        <div class="switch-link-row">
          <span class="prompt-text">還沒有帳號？</span>
          <button type="button" class="text-link-bold" @click="authMode = 'register'">
            註冊
          </button>
        </div>
      </div>
    </div>

    <!-- SCREEN 13: 建立帳號 (Register) -->
    <div v-else-if="authMode === 'register'" class="auth-screen-card">
      <div class="brand-header">
        <h1 class="brand-title">我也要</h1>
        <h2 class="brand-action-title">建立帳號</h2>
        <p class="brand-sub">建立帳號，開始累積你的行動與成果。</p>
      </div>

      <div class="auth-fields-block register-fields">
        <div class="form-field-group">
          <label class="field-label">使用者名稱</label>
          <div class="field-input-box">
            <input 
              v-model="regUsername" 
              placeholder="3–30 字元，可用英數字、底線與句點" 
              class="auth-input"
            />
          </div>
        </div>

        <div class="form-field-group">
          <label class="field-label">電子郵件</label>
          <div class="field-input-box with-code-btn">
            <input 
              v-model="regEmail" 
              type="email" 
              placeholder="輸入電子郵件" 
              class="auth-input"
            />
            <button type="button" class="send-code-btn" @click="handleSendCode">
              {{ codeSent ? '驗證碼已發送' : '發送驗證碼' }}
            </button>
          </div>
        </div>

        <div class="form-field-group">
          <label class="field-label">驗證碼</label>
          <div class="field-input-box">
            <input 
              v-model="regCode" 
              placeholder="輸入驗證碼" 
              class="auth-input"
            />
          </div>
        </div>

        <div class="form-field-group">
          <label class="field-label">密碼</label>
          <div class="field-input-box">
            <input 
              v-model="regPassword" 
              type="password" 
              placeholder="設定密碼" 
              class="auth-input"
            />
          </div>
        </div>

        <div class="form-field-group">
          <label class="field-label">確認密碼</label>
          <div class="field-input-box">
            <input 
              v-model="regConfirmPassword" 
              type="password" 
              placeholder="再次輸入密碼" 
              class="auth-input"
              @keyup.enter="handleRegister"
            />
          </div>
        </div>

        <button type="button" class="btn-primary-auth" @click="handleRegister">
          註冊
        </button>

        <div class="switch-link-row">
          <span class="prompt-text">已經有帳號？</span>
          <button type="button" class="text-link-bold" @click="authMode = 'login'">
            返回登入
          </button>
        </div>
      </div>
    </div>

    <!-- SCREEN 14: 忘記密碼 (Forgot Password) -->
    <div v-else-if="authMode === 'forgot'" class="auth-screen-card">
      <div class="brand-header">
        <h1 class="brand-title">我也要</h1>
        <h2 class="brand-action-title">忘記密碼</h2>
        <p class="brand-sub">輸入註冊時使用的電子郵件，我們會寄送重設密碼連結。</p>
      </div>

      <div class="auth-fields-block">
        <div class="form-field-group">
          <label class="field-label">電子郵件</label>
          <div class="field-input-box">
            <input 
              v-model="forgotEmail" 
              type="email" 
              placeholder="輸入電子郵件" 
              class="auth-input"
              @keyup.enter="handleForgotSubmit"
            />
          </div>
        </div>

        <button type="button" class="btn-primary-auth" @click="handleForgotSubmit">
          寄送重設連結
        </button>

        <div class="forgot-notice-box">
          <p class="notice-line">請留意收件匣與垃圾郵件資料夾。</p>
          <p class="notice-line">收到郵件後，點擊連結設定新密碼，再返回登入。</p>
        </div>

        <div class="switch-link-row center-text">
          <button type="button" class="text-link-bold" @click="authMode = 'login'">
            返回登入
          </button>
        </div>
      </div>
    </div>
  </div>
</template>

<script setup>
import { ref } from 'vue'

const emit = defineEmits(['login-success'])

const authMode = ref('login')
const alertMessage = ref('')

// Login states
const loginEmail = ref('user@example.com')
const loginPassword = ref('123456')

// Register states
const regUsername = ref('')
const regEmail = ref('')
const regCode = ref('')
const regPassword = ref('')
const regConfirmPassword = ref('')
const codeSent = ref(false)

// Forgot states
const forgotEmail = ref('')

function showAlert(msg) {
  alertMessage.value = msg
  setTimeout(() => {
    alertMessage.value = ''
  }, 2500)
}

function handleLogin() {
  if (!loginEmail.value.trim() || !loginPassword.value) {
    showAlert('請輸入電子郵件與密碼')
    return
  }
  showAlert('登入成功！')
  setTimeout(() => {
    emit('login-success', { email: loginEmail.value, username: 'mine_account' })
  }, 600)
}

function handleSendCode() {
  if (!regEmail.value.trim()) {
    showAlert('請先輸入電子郵件')
    return
  }
  codeSent.value = true
  regCode.value = '849201'
  showAlert('驗證碼已發送至信箱！(測試驗證碼：849201)')
}

function handleRegister() {
  if (!regUsername.value.trim()) {
    showAlert('請輸入使用者名稱')
    return
  }
  if (!regEmail.value.trim()) {
    showAlert('請輸入電子郵件')
    return
  }
  if (!regPassword.value || regPassword.value !== regConfirmPassword.value) {
    showAlert('密碼不一致，請再次確認')
    return
  }
  showAlert('帳號建立成功！')
  setTimeout(() => {
    emit('login-success', { email: regEmail.value, username: regUsername.value })
  }, 600)
}

function handleForgotSubmit() {
  if (!forgotEmail.value.trim()) {
    showAlert('請輸入電子郵件')
    return
  }
  showAlert('重設連結已寄出至您的信箱！')
}
</script>

<style scoped>
.auth-view-container {
  width: 100%;
  display: flex;
  flex-direction: column;
  gap: 16px;
  box-sizing: border-box;
}

.auth-nav-switch {
  display: flex;
  justify-content: center;
  gap: 8px;
  padding: 6px 0;
}

.auth-mode-btn {
  padding: 6px 14px;
  border-radius: 16px;
  border: 1px solid #DBDBD9;
  background: #FFFFFF;
  color: #75757A;
  font-size: 12px;
  cursor: pointer;
  transition: all 0.15s ease;
}

.auth-mode-btn.active {
  background: #141417;
  color: #FFFFFF;
  border-color: #141417;
  font-weight: 700;
}

.auth-alert-box {
  background: #141417;
  color: #FFFFFF;
  padding: 10px 16px;
  border-radius: 8px;
  font-size: 13px;
  text-align: center;
  animation: fadeIn 0.2s ease;
}

@keyframes fadeIn {
  from { opacity: 0; transform: translateY(-4px); }
  to { opacity: 1; transform: translateY(0); }
}

.auth-screen-card {
  background: #FFFFFF;
  border-radius: 12px;
  border: 1px solid #DBDBD9;
  padding: 32px 24px;
  display: flex;
  flex-direction: column;
  gap: 24px;
  box-sizing: border-box;
}

.brand-header {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.brand-title {
  color: #1F1F1F;
  font-size: 30px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  line-height: 1.2;
  margin: 0;
}

.brand-action-title {
  color: #1F1F1F;
  font-size: 20px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  margin: 0;
}

.brand-sub {
  color: #757575;
  font-size: 14px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 400;
  line-height: 1.5;
  margin: 0;
}

.auth-fields-block {
  display: flex;
  flex-direction: column;
  gap: 16px;
}

.register-fields {
  gap: 12px;
}

.form-title {
  color: #1F1F1F;
  font-size: 20px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  margin: 0 0 4px 0;
}

.form-field-group {
  display: flex;
  flex-direction: column;
  gap: 6px;
}

.field-label {
  color: #1F1F1F;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
}

.field-input-box {
  width: 100%;
  height: 50px;
  padding: 0 14px;
  background: #FBFBFB;
  border-radius: 8px;
  outline: 1px solid #D1D1D1;
  outline-offset: -1px;
  display: flex;
  align-items: center;
  box-sizing: border-box;
}

.field-input-box:focus-within {
  outline-color: #141417;
}

.field-input-box.with-code-btn {
  padding-right: 7px;
}

.auth-input {
  flex: 1;
  border: none;
  outline: none;
  background: transparent;
  font-size: 14px;
  font-family: inherit;
  color: #1F1F1F;
}

.auth-input::placeholder {
  color: #949494;
}

.send-code-btn {
  padding: 4px 14px;
  background: #141417;
  border-radius: 18px;
  border: none;
  color: #FFFFFF;
  font-size: 13px;
  font-weight: 500;
  cursor: pointer;
  white-space: nowrap;
  transition: background 0.15s ease;
}

.send-code-btn:hover {
  background: #25252b;
}

.forgot-link-row {
  display: flex;
  justify-content: flex-end;
}

.text-link-btn {
  background: none;
  border: none;
  color: #6B6B6B;
  font-size: 13px;
  font-family: 'Noto Sans TC', sans-serif;
  cursor: pointer;
  padding: 0;
}

.text-link-btn:hover {
  color: #141417;
}

.btn-primary-auth {
  width: 100%;
  height: 48px;
  background: #141414;
  border-radius: 8px;
  border: none;
  color: #FFFFFF;
  font-size: 15px;
  font-family: 'Noto Sans TC', sans-serif;
  font-weight: 700;
  cursor: pointer;
  transition: all 0.15s ease;
  display: flex;
  align-items: center;
  justify-content: center;
}

.btn-primary-auth:hover {
  background: #26262b;
}

.switch-link-row {
  display: flex;
  align-items: center;
  justify-content: center;
  gap: 8px;
  margin-top: 4px;
}

.prompt-text {
  color: #4D4D4D;
  font-size: 14px;
}

.text-link-bold {
  background: none;
  border: none;
  color: #141417;
  font-size: 14px;
  font-weight: 700;
  font-family: 'Noto Sans TC', sans-serif;
  cursor: pointer;
  padding: 0;
}

.text-link-bold:hover {
  text-decoration: underline;
}

.forgot-notice-box {
  display: flex;
  flex-direction: column;
  gap: 6px;
  padding: 8px 0;
}

.notice-line {
  color: #757575;
  font-size: 13px;
  line-height: 1.5;
  margin: 0;
}

.center-text {
  justify-content: center;
}
</style>
