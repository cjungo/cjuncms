<template>
  <CJunCmsPageMainLayout class="setting-profile-page">
    <ElForm :labelWidth="100">
      <ElRow>
        <ElCol :span="6">
          <ElFormItem label="头像">
            <ImageInput @input="onSetAvatar" />
          </ElFormItem>
        </ElCol>
        <ElCol :span="18">
          <ElRow>
            <ElCol :span="8">
              <ElFormItem label="ID">
                <ElInput :value="profile?.id" readonly />
              </ElFormItem>
            </ElCol>
            <ElCol :span="8">
              <ElFormItem label="用户名">
                <ElInput v-model="profile.username" :readonly="isReadonly" />
              </ElFormItem>
            </ElCol>
            <ElCol :span="8">
              <ElFormItem label="工号">
                <ElInput v-model="profile.jobnumber" :readonly="isReadonly" />
              </ElFormItem>
            </ElCol>
          </ElRow>
          <ElRow>
            <ElCol :span="8">
              <ElFormItem label="昵称">
                <ElInput v-model="profile.nickname" :readonly="isReadonly" />
              </ElFormItem>
            </ElCol>
          </ElRow>
        </ElCol>
      </ElRow>
    </ElForm>
    <ElButton @click="onSave">保存</ElButton>
  </CJunCmsPageMainLayout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { getProfile, setProfile, type Profile } from "../../apis/login";

const isReadonly = ref(true);
const avatarFile = ref<File>();

const profile = ref<Profile>({
  id: 0,
  jobnumber: "",
  username: "",
  nickname: "",
  fullname: "",
  birthday: "",
  avatar_path: "",
  is_removed: 0,
});

const onSetAvatar = (file: File) => {
  avatarFile.value = file;
};

const onSave = async () => {
  if (avatarFile.value) {
    // TODO 上传
  }
  await setProfile(profile.value);
};

onBeforeMount(async () => {
  const response = await getProfile();
  profile.value = response.data;
});
</script>

<style lang="scss" scoped>
.setting-profile-page {
  display: flex;
  flex-direction: column;
}
</style>
