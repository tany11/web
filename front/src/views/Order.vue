<template>
    <div class="order-container">
        <div class="order-form">
            <!-- ダッシュボードの内容をここに追加 -->
            <h2>受注表</h2>
            <ul>
                <li>
                    <label for="store-code">店舗コード: </label>
                    <input type="text" id="store-code" v-model="storeCode" @input="filterStores">
                </li>
                <li>
                    <label for="store-name">店名　　: </label>
                    <select id="store-name" name="store-name" v-model.number="storeID"
                        :class="{ 'is-invalid': showValidation && !isStoreNameValid }">
                        <option value="">選択してください</option>
                        <option v-for="store in filteredStores" :key="store.id" :value="store.id">
                            {{ store.name }}
                        </option>
                    </select>
                    <span v-if="showValidation && !isStoreNameValid" class="error-message">店名を選択してください</span>
                </li>
                <li>
                    <label for="phone-number">電話番号: </label>
                    <input type="text" id="phone-number" name="phone-number" v-model="phoneNumber"
                        @blur="fetchCustomerInfo" :class="{ 'is-invalid': showValidation && !isPhoneNumberValid }">
                    <a href="#" @click.prevent="showCustomerDetail" v-if="phoneNumber">詳細</a>
                    <span v-if="showValidation && !isPhoneNumberValid" class="error-message">有効な電話番号を入力してください</span>
                </li>
                <li>
                    <label for="customer-name">お客様名: </label>
                    <input type="text" id="customer-name" name="customer-name" v-model="customerName">
                </li>
                <li>
                    <label for="model-name">モデル名: </label>
                    <input type="text" id="model-name" name="model-name" v-model="modelName">
                </li>
                <li>
                    <label for="actual-model">実モデル: </label>
                    <select id="actual-model" name="actual-model" v-model="actualModel">
                        <option value="">選択してください</option>
                        <option v-for="cast in castList" :key="cast.cast_id" :value="cast.cast_id">
                            {{ cast.name }}
                        </option>
                    </select>
                </li>
                <li class="course-item">
                    <label for="course-name">コース　: </label>
                    <div class="course-inputs">
                        <input type="text" id="course-name" name="course-name" v-model="courseMin">
                        <span>+</span>
                        <input type="text" id="extra-course" name="extra-course" v-model="extraCourse">
                        <span>分</span>
                    </div>
                </li>
                <li>
                    <label for="price">料金　　: </label>
                    <input type="number" id="price" name="price" v-model.number="price" step="1000" min="0">
                </li>
                <!-- 自宅orホテルのプルダウンもほしいか -->
                <li>
                    <label for="postal-code">郵便番号: </label>
                    <input type="text" id="postal-code" name="postal-code" v-model="postalCode" @blur="fetchAddress">
                </li>
                <!-- 住所を使用してHomesとGoogleMapのURLもほしい -->
                <li>
                    <label for="address">住所　　: </label>
                    <input type="text" id="address" name="address" v-model="address">
                </li>
                <li>
                    <label for="delivery">送り　　: </label>
                    <select id="delivery" name="delivery" v-model="driverID">
                        <option value="">選択してください</option>
                        <option v-for="staff in driverStaffList" :key="staff.id" :value="staff.staff_id">
                            {{ staff.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <label for="nomination-fee">指名料　: </label>
                    <input type="number" id="nomination-fee" name="nomination-fee" v-model.number="reservationFee"
                        step="1000" min="0">
                </li>
                <li>
                    <label for="transportation-fee">交通費　: </label>
                    <input type="number" id="transportation-fee" name="transportation-fee"
                        v-model.number="transportationFee" step="1000" min="0">
                </li>
                <li>
                    <label for="travel-expenses">出張費　: </label>
                    <input type="number" id="travel-expenses" name="travel-expenses" v-model.number="travelCost"
                        step="1000" min="0">
                </li>
                <li>
                    <label for="media">媒体　　: </label>
                    <select id="media" name="media" v-model.number="media">
                        <option value="">選択してください</option>
                        <option v-for="media in mediaList" :key="media.id" :value="media.id">
                            {{ media.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <label for="notes">備考　　: </label>
                    <textarea id="notes" name="notes" v-model="notes" rows="3"></textarea>
                </li>
                <li>
                    <label for="card-handler">カード　: </label>
                    <select id="card-handler" name="card-handler" v-model="cardstaffID">
                        <option v-for="staff in officeStaffList" :key="staff.id" :value="staff.staff_id">
                            {{ staff.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <label for="order-taker">受注者　: </label>
                    <select id="order-taker" name="order-taker" v-model="orderStaffID">
                        <option v-for="staff in officeStaffList" :key="staff.id" :value="staff.staff_id">
                            {{ staff.name }}
                        </option>
                    </select>
                </li>
                <li>
                    <button @click="submitOrder">受注</button>
                </li>
            </ul>
        </div>
        <div class="order-display">
            <div v-if="loading">データを読み込んでいます...</div>
            <div v-else-if="orders.length === 0">
                オーダーデータがありません。
                <button @click="toggleShowCompleted" class="completed-toggle">
                    {{ showCompleted ? '未確定のみ表示' : '確定済を表示' }}
                </button>
            </div>
            <div v-else>
                <div class="pagination">
                    <span>{{ currentPage }}/{{ totalPages }}</span>
                    <button @click="prevPage" :disabled="currentPage === 1 || isEditing">前へ</button>
                    <button @click="nextPage" :disabled="currentPage === totalPages || isEditing">次へ</button>
                </div>
                <div v-if="currentOrder" class="order-box">
                    <div class="order-header">
                        <h3>オーダー明細</h3>
                        <button @click="toggleShowCompleted" class="completed-toggle">
                            {{ showCompleted ? '確定済を非表示' : '確定済を表示' }}
                        </button>
                    </div>
                    <form @submit.prevent="updateOrder">
                        <template v-for="(value, key) in displayableFields" :key="key">
                            <p>
                                <strong>{{ getFieldLabel(key) }}: </strong>
                                <template v-if="isEditing && isEditableField(key)">
                                    <template v-if="key === 'ActualModel'">
                                        <select v-model="editedOrder[key]"
                                            :class="{ 'is-invalid': editValidation && !isEditedActualModelValid }">
                                            <option v-for="cast in castList" :key="cast.cast_id" :value="cast.cast_id">
                                                {{ cast.name }}
                                            </option>
                                        </select>
                                        <span v-if="editValidation && !isEditedActualModelValid"
                                            class="error-message">実モデルを選択してください</span>
                                    </template>
                                    <template v-else-if="key === 'StoreID'">
                                        <select v-model="editedOrder[key]"
                                            :class="{ 'is-invalid': editValidation && !isEditedStoreIDValid }">
                                            <option v-for="store in storeList" :key="store.id" :value="store.id">
                                                {{ store.name }}
                                            </option>
                                        </select>
                                        <span v-if="editValidation && !isEditedStoreIDValid"
                                            class="error-message">店舗を選択してください</span>
                                    </template>
                                    <template v-else-if="key === 'Media'">
                                        <select v-model="editedOrder[key]">
                                            <option v-for="media in mediaList" :key="media.id" :value="media.id">
                                                {{ media.name }}
                                            </option>
                                        </select>
                                    </template>
                                    <template
                                        v-else-if="key === 'CardstaffID' || key === 'OrderStaffID' || key === 'DriverID'">
                                        <select v-model="editedOrder[key]"
                                            :class="{ 'is-invalid': editValidation && key === 'DriverID' && !isEditedDriverIDValid }">
                                            <option v-for="staff in getStaffList(key)" :key="staff.staff_id"
                                                :value="staff.staff_id">
                                                {{ staff.name }}
                                            </option>
                                        </select>
                                        <span v-if="editValidation && key === 'DriverID' && !isEditedDriverIDValid"
                                            class="error-message">ドライバーを選択してください</span>
                                    </template>
                                    <template v-else-if="isPriceField(key) || isIntField(key)">
                                        <input v-model.number="editedOrder[key]" type="number" :step="getStepValue(key)"
                                            min="0" @input="validateIntInput(key)"
                                            :class="{ 'is-invalid': editValidation && ((key === 'Price' && !isEditedPriceValid) || (key === 'CourseMin' && !isEditedCourseMinValid)) }">
                                        <span v-if="editValidation && key === 'Price' && !isEditedPriceValid"
                                            class="error-message">有効な料金を入力してください</span>
                                        <span v-if="editValidation && key === 'CourseMin' && !isEditedCourseMinValid"
                                            class="error-message">有効なコース時間を入力してください</span>
                                    </template>
                                    <template v-else-if="key === 'Notes'">
                                        <textarea v-model="editedOrder[key]" rows="3"></textarea>
                                    </template>
                                    <template v-else>
                                        <input v-model="editedOrder[key]" :type="getInputType(key)"
                                            :class="{ 'is-invalid': editValidation && ((key === 'CustomerName' && !isEditedCustomerNameValid) || (key === 'PhoneNumber' && !isEditedPhoneNumberValid) || (key === 'Address' && !isEditedAddressValid)) }">
                                        <span
                                            v-if="editValidation && key === 'CustomerName' && !isEditedCustomerNameValid"
                                            class="error-message">お客様名を入力してください</span>
                                        <span
                                            v-if="editValidation && key === 'PhoneNumber' && !isEditedPhoneNumberValid"
                                            class="error-message">有効な電話番号を入力してください</span>
                                        <span v-if="editValidation && key === 'Address' && !isEditedAddressValid"
                                            class="error-message">住所を入力してください</span>
                                    </template>
                                </template>
                                <template v-else>
                                    {{ getDisplayValue(key, value) }}
                                </template>
                            </p>
                        </template>
                        <div class="button-group">
                            <button type="button" @click="toggleEdit" v-if="!isEditing">編集</button>
                            <button type="button" @click="saveEdit" v-if="isEditing">保存</button>
                            <button type="button" @click="cancelEdit" v-if="isEditing">キャンセル</button>
                            <button type="button" @click="confirmOrder" :disabled="isEditing">確定</button>
                        </div>
                    </form>
                    <button type="button" @click="confirmDelete" class="delete-button" :disabled="isEditing">削除</button>
                </div>
                <div v-else>
                    選択されたオーダーがありません。
                </div>
            </div>
        </div>
        <CustomerDetail v-if="showCustomerModal" :customer="customerDetail" @close="closeCustomerModal" />
    </div>
</template>

<script>
import axios from 'axios';
import { mapState } from 'vuex';
import CustomerDetail from '@/components/CustomerDetail.vue';

export default {
    name: 'Order',
    components: {
        CustomerDetail
    },
    data() {
        return {
            phoneNumber: '',
            customerName: '',
            address: '',
            storeID: null,
            modelName: '',
            actualModel: '',
            courseMin: '',
            extraCourse: '',
            price: 0,
            postalCode: '',
            reservationFee: 0,
            transportationFee: 0,
            travelCost: 0,
            media: null,
            notes: '',
            orders: [],
            currentPage: 1,
            totalPages: 1,
            loading: true,
            driverID: '',
            cardstaffID: '',
            orderStaffID: '',
            staffList: [],
            castList: [],
            isEditing: false,
            editedOrder: {},
            showValidation: false,
            storeList: [],
            mediaList: [],
            showCompleted: false,
            showCustomerModal: false,
            customerDetail: null,
            storeCode: '',
            filteredStores: [],
            editValidation: false,
        }
    },
    computed: {
        currentOrder() {
            if (!this.orders || this.orders.length === 0) return null;
            return this.orders[this.currentPage - 1] || null;
        },
        driverStaffList() {
            return this.staffList.filter(staff => staff.driver_flg === "1");
        },
        officeStaffList() {
            return this.staffList.filter(staff => staff.office_flg === "1");
        },
        displayableFields() {
            const excludedFields = ['GroupID', 'CustomerID', 'UpdatedAt', 'CompletionFlg', 'IsDeleted', 'ExtraCourse', 'PostalCode'];
            return Object.keys(this.currentOrder || {}).reduce((acc, key) => {
                if (!excludedFields.includes(key)) {
                    acc[key] = this.currentOrder[key];
                }
                return acc;
            }, {});
        },
        isPhoneNumberValid() {
            // 日本の電話番号の簡単なバリデーション
            const phoneRegex = /^(0\d{1,4}-?\d{1,4}-?\d{4})$/;
            return phoneRegex.test(this.phoneNumber);
        },
        isStoreIDValid() {
            return this.storeID !== null && this.storeID !== '';
        },
        isEditedOrderValid() {
            return this.isEditedStoreIDValid &&
                this.isEditedCustomerNameValid &&
                this.isEditedActualModelValid &&
                this.isEditedCourseMinValid &&
                this.isEditedPriceValid &&
                this.isEditedAddressValid &&
                this.isEditedDriverIDValid &&
                this.isEditedPhoneNumberValid;
        },
        isEditedStoreIDValid() {
            return this.editedOrder.StoreID !== null && this.editedOrder.StoreID !== '';
        },
        isEditedCustomerNameValid() {
            return this.editedOrder.CustomerName && this.editedOrder.CustomerName.trim() !== '';
        },
        isEditedActualModelValid() {
            return this.editedOrder.ActualModel !== '';
        },
        isEditedCourseMinValid() {
            const courseMinInt = parseInt(this.editedOrder.CourseMin, 10);
            return !isNaN(courseMinInt) && courseMinInt > 0;
        },
        isEditedPriceValid() {
            return this.editedOrder.Price !== '' && this.editedOrder.Price > 0;
        },
        isEditedAddressValid() {
            return this.editedOrder.Address && this.editedOrder.Address.trim() !== '';
        },
        isEditedDriverIDValid() {
            return this.editedOrder.DriverID !== '';
        },
        isEditedPhoneNumberValid() {
            const phoneRegex = /^(0\d{1,4}-?\d{1,4}-?\d{4})$/;
            return phoneRegex.test(this.editedOrder.PhoneNumber);
        },
        isFormValid() {
            return this.isStoreIDValid &&
                this.isPhoneNumberValid;
        },
        ...mapState(['apiBaseUrl']),

    },
    methods: {
        async fetchCustomerInfo() {
            if (this.phoneNumber) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/phone/${this.phoneNumber}`);
                    if (response.data && response.data.data) {
                        this.$nextTick(() => {
                            this.customerName = response.data.data.CustomerName || '';
                            this.address = response.data.data.Address || '';
                            console.log('データが更新されました:', this.customerName, this.address);
                        });
                    }
                } catch (error) {
                    console.error('顧客情報の取得に失敗しました:', error);
                }
            }
        },
        async fetchAddress() {
            if (this.postalCode) {
                try {
                    const response = await axios.get(`https://zipcloud.ibsnet.co.jp/api/search?zipcode=${this.postalCode}`);
                    if (response.data && response.data.results) {
                        const result = response.data.results[0];
                        this.address = `${result.address1}${result.address2}${result.address3}`;
                    } else {
                        console.log('該当する住所が見つかりませんでした');
                    }
                } catch (error) {
                    console.error('住所情報の取得に失敗しました:', error);
                }
            }
        },
        validateForm() {
            this.showValidation = true;
            if (!this.isFormValid) {
                alert('フォームに無効な入力があります。すべての必須フィールドを正しく入力してください。');
                return false;
            }
            return true;
        },
        async submitOrder() {
            if (!this.validateForm()) {
                return;
            }

            try {
                const orderData = {
                    storeID: this.storeID,
                    phoneNumber: this.phoneNumber,
                    customerName: this.customerName,
                    modelName: this.modelName,
                    actualModel: this.actualModel,
                    courseMin: parseInt(this.courseMin, 10),
                    extraCourse: parseInt(this.extraCourse, 10) || 0,
                    price: parseInt(this.price, 10),
                    postalCode: this.postalCode,
                    address: this.address,
                    driverID: this.driverID,
                    reservationFee: parseInt(this.reservationFee, 10),
                    transportationFee: parseInt(this.transportationFee, 10),
                    travelCost: parseInt(this.travelCost, 10),
                    media: this.media,
                    notes: this.notes,
                    cardstaffID: this.cardstaffID,
                    orderStaffID: this.orderStaffID,
                };

                const response = await axios.post(`${this.apiBaseUrl}/orders`, orderData);
                console.log('注文が正常に送信されました:', response.data);

                // 成功メッセージを表示
                alert('注文が正常に送信されました。');

                // フォームをリセットし、オーダー一覧を更新
                this.resetForm();
                await this.fetchOrders();

                // 最新のオーダーを表示するために currentPage を 1 にセット
                this.currentPage = 1;
            } catch (error) {
                console.error('注文の送信に失敗しました:', error);
                alert('注文の送信に失敗しました。もう一度お試しください。');
            }
        },
        resetForm() {
            // フォームの各フィールドをリセット
            this.phoneNumber = '';
            this.customerName = '';
            this.address = '';
            this.storeID = null;
            this.modelName = '';
            this.actualModel = '';
            this.courseMin = '';
            this.extraCourse = '';
            this.price = '';
            this.postalCode = '';
            this.transportationFee = '';
            this.media = null;
            this.notes = '';
            this.driverID = '';
            this.cardstaffID = '';
            this.orderStaffID = '';
            this.reservationFee = '';
            this.travelCost = '';
            this.showValidation = false;
        },
        async fetchOrders() {
            this.loading = true;
            try {
                const endpoint = this.showCompleted ? '/orders' : '/orders/reserved';
                const response = await axios.get(`${this.apiBaseUrl}${endpoint}`);
                this.orders = response.data.data || []; // データが data プロパティ内にある場合
                this.totalPages = this.orders.length;
                this.currentPage = this.orders.length > 0 ? 1 : 0; // 最新のオーダーを表示するために1ページ目にセット
                console.log('Fetched orders:', this.orders);
                console.log('Current page after fetch:', this.currentPage);
            } catch (error) {
                console.error('オーダーの取得に失敗しました:', error);
                this.orders = [];
                this.totalPages = 0;
                this.currentPage = 0;
            } finally {
                this.loading = false;
            }
        },
        prevPage() {
            if (this.currentPage > 1 && !this.isEditing) {
                this.currentPage--;
            }
        },
        nextPage() {
            if (this.currentPage < this.totalPages && !this.isEditing) {
                this.currentPage++;
            }
        },

        //プルダウンリストを取得
        async fetchStaffList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/staff/dropdown`);
                this.staffList = response.data.data || [];
            } catch (error) {
                console.error('スタッフリストの取得に失敗しました:', error);
            }
        },
        async fetchCastList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/cast/dropdown`);
                this.castList = response.data.data || [];
            } catch (error) {
                console.error('キャストリストの取得に失敗しました:', error);
            }
        },
        async fetchStoreList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/store/dropdown`);
                this.storeList = response.data.data || [];
                this.filteredStores = [...this.storeList]; // 初期状態では全ての店舗を表示
                console.log('店舗リスト:', this.storeList);
            } catch (error) {
                console.error('店舗リストの取得に失敗しました:', error);
            }
        },
        async fetchMediaList() {
            try {
                const response = await axios.get(`${this.apiBaseUrl}/media/dropdown`);
                this.mediaList = response.data.data || [];
                console.log('媒体リスト:', this.mediaList);
            } catch (error) {
                console.error('媒体リストの取得に失敗しました:', error);
            }
        },

        // 明細で使用するリストを取得
        getCastName(castId) {
            const cast = this.castList.find(c => c.cast_id === castId);
            return cast ? cast.name : 'Unknown';
        },
        getStaffName(staffId) {
            const staff = this.staffList.find(s => s.staff_id === staffId);
            return staff ? staff.name : 'Unknown';
        },
        getStoreName(storeId) {
            const store = this.storeList.find(s => s.id === storeId);
            return store ? store.name : 'Unknown';
        },
        getMediaName(mediaId) {
            const media = this.mediaList.find(m => m.id === mediaId);
            return media ? media.name : 'Unknown';
        },
        toggleEdit() {
            if (!this.isEditing) {
                this.editedOrder = { ...this.currentOrder };
            }
            this.isEditing = !this.isEditing;
        },
        saveEdit() {
            this.editValidation = true;
            if (this.isEditedOrderValid) {
                this.updateOrder();
            } else {
                alert('フォームに無効な入力があります。すべての必須フィールドを正しく入力してください。');
            }
        },
        cancelEdit() {
            this.editedOrder = { ...this.currentOrder };
            this.isEditing = false;
            this.editValidation = false;
        },
        async updateOrder() {
            if (!this.currentOrder) {
                console.error('現在のオーダーが見つかりません');
                console.error(this.currentOrder);
                return;
            }

            const orderId = this.currentOrder.ID;
            if (!orderId) {
                console.error('現在のオーダーIDが見つかりません');
                console.error(orderId);
                return;
            }

            // INT項目の値を整数に変換
            const intFields = ['CourseMin', 'ExtraTime', 'ExtraCourse', 'Price', 'ReservationFee', 'TransportationFee', 'TravelCost'];
            intFields.forEach(field => {
                if (this.editedOrder[field] !== undefined) {
                    this.editedOrder[field] = Math.floor(Number(this.editedOrder[field]));
                }
            });

            try {
                await axios.put(`${this.apiBaseUrl}/orders/${orderId}`, this.editedOrder);
                this.fetchOrders();
                this.isEditing = false;
            } catch (error) {
                console.error('オーダーの更新に失敗しました:', error);
            }
        },
        async confirmDelete() {
            if (this.isEditing) return; // 編集中は処理を実行しない

            if (confirm('このオーダーを削除しても��ろしいですか？')) {
                try {
                    await axios.put(`${this.apiBaseUrl}/orders/${this.currentOrder.ID}/delete`);
                    this.fetchOrders();
                } catch (error) {
                    console.error('オーダーの削除に失敗しました:', error);
                }
            }
        },
        async confirmOrder() {
            if (this.isEditing) return; // 編集中は処理を実行しない

            if (confirm('このオーダーを確定してもよろしいですか？')) {
                try {
                    await axios.put(`${this.apiBaseUrl}/orders/${this.currentOrder.ID}/completion`);
                    this.fetchOrders();
                } catch (error) {
                    console.error('オーダーの確定に失敗しまた:', error);
                }
            }
        },
        isEditableField(key) {
            const editableFields = [
                'CustomerName',
                'PhoneNumber',
                'ModelName',
                'ActualModel',
                'CourseMin',
                'ExtraTime',
                'Price',
                'Address',
                'ReservationFee',
                'TransportationFee',
                'TravelCost',
                'Notes',
                'CardstaffID',
                'DriverID',
                'StoreID',
                'Media',
                'OrderStaffID'
            ];
            return editableFields.includes(key);
        },
        isPriceField(key) {
            return ['Price', 'ReservationFee', 'TransportationFee', 'TravelCost'].includes(key);
        },
        isIntField(key) {
            return ['CourseMin', 'ExtraTime', 'ExtraCourse'].includes(key);
        },
        getStepValue(key) {
            if (this.isPriceField(key)) {
                return 1000;
            } else if (key === 'CourseMin') {
                return 10;
            } else if (key === 'ExtraTime' || key === 'ExtraCourse') {
                return 30;
            }
            return 1;
        },
        validateIntInput(key) {
            if (this.isIntField(key)) {
                this.editedOrder[key] = Math.floor(this.editedOrder[key]);
            }
        },
        getInputType(key) {
            return this.isPriceField(key) ? 'number' : 'text';
        },
        getStaffList(key) {
            if (key === 'DriverID') {
                return this.driverStaffList;
            } else if (key === 'CardstaffID' || key === 'OrderStaffID') {
                return this.officeStaffList;
            }
            return [];
        },
        getFieldLabel(key) {
            const labels = {
                StoreID: '店名　　　',
                CustomerName: 'お客様名　',
                PhoneNumber: '電話番号　',
                ModelName: 'モデル名　',
                ActualModel: '実モデル　',
                CourseMin: 'コース　　',
                ExtraTime: '延長時間　',
                Price: '料金　　　',
                Address: '住所　　　',
                DriverID: '送り　　　',
                ReservationFee: '指名料　　',
                TransportationFee: '交通費　　',
                TravelCost: '出張費　　',
                Media: '媒体　　　',
                Notes: '備考　　　',
                CardstaffID: 'カード　　',
                OrderStaffID: '受注者　　',
                CreatedAt: '受注日　　',
            };
            return labels[key] || key;
        },
        formatDate(dateString) {
            if (!dateString) return '';
            const date = new Date(dateString);
            const japanTime = new Date(date.toLocaleString("en-US", { timeZone: "Asia/Tokyo" }));

            const year = String(japanTime.getFullYear()).slice(-2);
            const month = String(japanTime.getMonth() + 1).padStart(2, '0');
            const day = String(japanTime.getDate()).padStart(2, '0');
            const hours = String(japanTime.getHours()).padStart(2, '0');
            const minutes = String(japanTime.getMinutes()).padStart(2, '0');

            return `${year}/${month}/${day} ${hours}:${minutes}`;
        },
        getDisplayValue(key, value) {
            if (key === 'ActualModel') {
                return this.getCastName(value);
            } else if (key === 'DriverID' || key === 'CardstaffID' || key === 'OrderStaffID') {
                return this.getStaffName(value);
            } else if (key === 'StoreID') {
                return this.getStoreName(value);
            } else if (key === 'Media') {
                return this.getMediaName(value);
            } else if (this.isPriceField(key)) {
                return `¥${value.toLocaleString()}`;
            } else if (key === 'Notes') {
                return value || '';
            } else if (key === 'CreatedAt') {
                return this.formatDate(value);
            } else if (key === 'CourseMin') {
                const extraCourse = this.currentOrder.ExtraCourse || 0;
                return `${value}分 + ${extraCourse}分`;
            } else if (key === 'ExtraTime') {
                return `${value}分`;
            } else if (key === 'PhoneNumber') {
                return `184${value}`;
            }
            return value;
        },
        async toggleShowCompleted() {
            this.showCompleted = !this.showCompleted;
            await this.fetchOrders();
            if (this.orders.length === 0) {
                this.currentPage = 0;
                this.totalPages = 0;
            } else {
                this.currentPage = 1;
                this.totalPages = this.orders.length;
            }
        },
        async showCustomerDetail() {
            if (this.phoneNumber) {
                try {
                    const response = await axios.get(`${this.apiBaseUrl}/customers/detail/${this.phoneNumber}`);
                    if (response.data && response.data.data) {
                        this.customerDetail = response.data.data;
                        this.showCustomerModal = true;
                    }
                } catch (error) {
                    console.error('顧客詳細の取得に失敗しました:', error);
                }
            }
        },
        closeCustomerModal() {
            this.showCustomerModal = false;
        },
        filterStores() {
            if (this.storeCode === '') {
                this.filteredStores = [...this.storeList];
            } else {
                this.filteredStores = this.storeList.filter(store =>
                    store.store_code.toString().startsWith(this.storeCode)
                );
            }

            // フィルタリング後に店舗が1つだけ残った場合、自動的にその店舗を選択
            if (this.filteredStores.length === 1) {
                this.storeID = this.filteredStores[0].id;
            } else {
                this.storeID = null; // 複数の選択肢がある場合はクリア
            }
        },
    },
    mounted() {
        this.fetchOrders();
        this.fetchStaffList();
        this.fetchCastList();
        this.fetchStoreList();
        this.fetchMediaList();
    }
}

</script>

<style scoped>
.order-box {
    position: relative;
    padding-bottom: 50px;
    /* 削除ボタンの高さ分の余白を追加 */
    border: 1px solid var(--color-border);
    border-radius: 5px;
    padding: 15px;
    margin-bottom: 20px;
}

.button-group {
    margin-top: 20px;
}

.button-group button {
    margin-right: 10px;
}

.delete-button {
    position: absolute;
    bottom: 15px;
    right: 15px;
    background-color: #ff4d4d;
    /* 赤色の背景 */
    color: var(--vt-c-white);
    border: none;
    padding: 8px 15px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.delete-button:hover {
    background-color: #ff1a1a;
    /* ホバー時の色 */
}

@media (prefers-color-scheme: dark) {
    .order-box {
        border-color: var(--color-border);
    }

    .delete-button {
        background-color: #8b0000;
        /* ダークモード時の赤色 */
    }

    .delete-button:hover {
        background-color: #a50000;
        /* ダークモード時のホバー色 */
    }
}

.is-invalid {
    border-color: red;
}

.error-message {
    color: red;
    font-size: 0.8em;
}


@media (max-width: 768px) {
    .order-container {
        padding-left: 0;
        padding-bottom: 60px;
        /* モバイル時のサイドバーの高さに応じて調整 */
    }
}

.order-header {
    display: flex;
    justify-content: space-between;
    align-items: center;
    margin-bottom: 15px;
}

.completed-toggle {
    background-color: #4CAF50;
    color: white;
    border: none;
    padding: 8px 15px;
    border-radius: 4px;
    cursor: pointer;
    transition: background-color 0.3s;
}

.completed-toggle:hover {
    background-color: #45a049;
}

@media (prefers-color-scheme: dark) {
    .completed-toggle {
        background-color: #2E7D32;
    }

    .completed-toggle:hover {
        background-color: #1B5E20;
    }
}

.course-inputs {
    display: flex;
    align-items: center;
}

.course-inputs input[type="text"] {
    width: 50px;
    margin-right: 5px;
}

.course-inputs span {
    margin: 0 5px;
}

/* test */
.course-item {
    display: flex;
    align-items: center;
    margin-bottom: 10px;
    list-style: disc;
    /* ーカーを強制的に表示 */
    list-style-position: outside;
    /* マーカーをリストの外側に配置 */
}

.course-item label {
    margin-right: 10px;
    /* ラベルと入力フィールドの間に余白を追加 */
    white-space: nowrap;
    /* ラベルが折り返されないようにする */
}

.pagination button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.button-group button:disabled,
.delete-button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.course-inputs {
    display: flex;
    align-items: center;
}

.pagination button:disabled,
.delete-button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.pagination button:disabled {
    opacity: 0.5;
    cursor: not-allowed;
}

.course-inputs input {
    margin-right: 5px;
    /* 各入力フィールド間の余白 */
}

.error-message {
    color: red;
    margin-top: 5px;
    display: block;
    /* エラーメッセージが新しい行に表示されるようにする */
}

.completed-toggle {
    margin-top: 10px;
}

.store-input {
    display: flex;
    align-items: center;
}

.store-input input[type="text"] {
    width: 80px;
    margin-right: 10px;
}
</style>