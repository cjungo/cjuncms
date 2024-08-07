<template>
  <CJunCmsPageMainLayout class="setting-profile-page">
    <ElForm labelWidth="5em" class="setting-profile-form">
      <ElRow>
        <ElCol :xs="6" :sm="6" :md="4" :lg="2" :xl="2">
          <ElFormItem label="头像">
            <ImageInput v-model="profile.avatar_path" @change="onSetAvatar" />
          </ElFormItem>
        </ElCol>
        <ElCol :xs="18" :sm="18" :md="18" :lg="12" :xl="12">
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
    <template #list>
      <div class="setting-profile-operation-bar">
        <ElButton v-if="isReadonly" @click="onClickEdit">编辑</ElButton>
        <template v-else>
          <ElButton @click="onClickCancel">取消</ElButton>
          <ElButton @click="onClickSave">保存</ElButton>
        </template>
      </div>
    </template>
  </CJunCmsPageMainLayout>
</template>

<script lang="ts" setup>
import { onBeforeMount, ref } from "vue";
import { getProfile, setProfile, type Profile } from "../../apis/login";
import { uploadAvatar } from "../../apis/storage";
import { getSuffix } from "../../utils/file";
import { cloneDeep } from "lodash";

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
const profileStore = ref<Profile>(profile.value);

const onSetAvatar = (file: File) => {
  avatarFile.value = file;
};

const onClickEdit = () => {
  isReadonly.value = false;
};

const onClickCancel = () => {
  isReadonly.value = true;
  profile.value = cloneDeep(profileStore.value);
};

const onClickSave = async () => {
  if (avatarFile.value) {
    const now = new Date().getTime();
    const suffix = getSuffix(avatarFile.value.name, true);
    const avatarResponse = await uploadAvatar(
      avatarFile.value,
      `${profile.value.id}-${now}${suffix}`
    );
    profile.value.avatar_path = avatarResponse.data;
  }
  await setProfile(profile.value);
};

onBeforeMount(async () => {
  const response = await getProfile();
  profile.value = response.data;
  profileStore.value = cloneDeep(response.data);
});
</script>

<style lang="scss" scoped>
.setting-profile-page {
  display: flex;
  flex-direction: column;

  background-color: #f6f8f9;
}

.setting-profile-form {
  padding: 1em;

  background-color: #fff;
  border: 1px solid #d8d9df;
  border-radius: .5em;
}

.setting-profile-operation-bar {
  display: flex;
  flex-direction: row;
  padding: 1em;
  
  background-color: #fff;
  border: 1px solid #d8d9df;
  border-radius: .5em;
}
</style>
