<template>
<div class="button">
  <v-btn v-if="user_info.isUser" class="btn mr-2" dark color="indigo" href="">My page</v-btn>
  <v-btn v-else class="btn mr-2" dark color="indigo" href="">Log in</v-btn>
</div>
</template>