<script>
    import { Input } from "sveltestrap";
    
    import Panel from "../../components/Panel.svelte";
    import Input_custom from '../../components/InputCustom.svelte' 
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
	export let listBranch = []
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "WAREHOUSE"
    let sData = "";
    let myModal_newentry = "";
    let flag_id_field = false;
    let flag_idbranch_field = false;
    let flag_btnsave = true;
    let code_branch_display = ""
    let warehouse_idbranch_select = "";
    let warehouse_idbranch_field = "";
    let warehouse_name_field = "";
    let warehouse_alamat_field = "";
    let warehouse_phone1_field = "";
    let warehouse_phone2_field = "";
    let warehouse_status_field = "";
    let warehouse_create_field = "";
    let warehouse_update_field = "";

    //==STORAGE==
    let sDataStorage = "";
    let searchStorage = "";
    let filterStorage = [];
    let total_liststorage = 0;
    let liststorage = [];
    let storage_warehouse_id = "";
    let storage_warehouse_title = "";
    let storage_flag_id = false;
    let storage_id_field = "";
    let storage_name_field = "";
    let storage_status_field = "";
    let storage_create_field = "";
    let storage_update_field = "";

    //==BIN==
    let sDataBin = "";
    let total_listbin = 0;
    let listuom = [];
    let listbin = [];
    let bin_storage_id = "";
    let bin_flag_id = false;
    let bin_id_field = 0;
    let bin_idstorage_field = "";
    let bin_iduom_field = "";
    let bin_name_field = "";
    let bin_maxcapacity_field = 0;
    let bin_status_field = "";
    let bin_create_field = "";
    let bin_update_field = "";


    let idrecord = "";
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    
    
    const NewData = (e,id,idbranch,name,alamat,phone1,phone2,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
        }else{
            flag_idbranch_field = true;
            flag_id_field = true;
            idrecord = id
            warehouse_idbranch_field = idbranch;
            warehouse_name_field = name;
            warehouse_alamat_field = alamat;
            warehouse_phone1_field = phone1;
            warehouse_phone2_field = phone2;
            warehouse_status_field = status;
            warehouse_create_field = create;
            warehouse_update_field = update;
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
        
    };
    const showStorage = (idwarehouse) => {
        storage_warehouse_id = idwarehouse
        storage_warehouse_title = idwarehouse
        call_storage(idwarehouse)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalliststorage"));
        myModal_newentry.show();
    };
    const showStorageBin = (idstorage) => {
        bin_storage_id = idstorage
        bin_idstorage_field = idstorage
        call_bin(idstorage)
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalliststoragebin"));
        myModal_newentry.show();
    };
    const call_formliststorage = (e,idstorage,nmstorage,statustorage,create,update) => {
        sDataStorage = e
        storage_flag_id = false;
        if(sDataStorage == "Edit"){
            storage_flag_id = true;
            storage_id_field = idstorage;
            storage_name_field = nmstorage;
            storage_status_field = statustorage;
            storage_create_field = create;
            storage_update_field = update;
        }

        myModal_newentry = new bootstrap.Modal(document.getElementById("modalcrudstorage"));
        myModal_newentry.show();
        
    };
    const call_formlistbin = (e,idbin,iduom,nmbin,max,status,create,update) => {
        call_uom()
        sDataBin = e
        if(sDataBin == "Edit"){
            storage_flag_id = true;
            bin_id_field = idbin;
            bin_iduom_field = iduom;
            bin_name_field = nmbin;
            bin_maxcapacity_field = max;
            bin_status_field = status;
            bin_create_field = create;
            bin_update_field = update;
        }

        myModal_newentry = new bootstrap.Modal(document.getElementById("modalcrudbin"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    
    
    async function handleSave() {
        let flag = true
        let msg = ""
        if(sData == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The ID is required\n"
            }
            if(idrecord.length == 10){
                flag = false
                msg += "The CODE is maxlength 5\n"
            }
            if(warehouse_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(warehouse_alamat_field == ""){
                flag = false
                msg += "The Alamat is required\n"
            }
            if(warehouse_phone1_field == ""){
                flag = false
                msg += "The Phone1 is required\n"
            }
            if(warehouse_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The CODE is required\n"
            }
            if(warehouse_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(warehouse_alamat_field == ""){
                flag = false
                msg += "The Alamat is required\n"
            }
            if(warehouse_phone1_field == ""){
                flag = false
                msg += "The Phone1 is required\n"
            }
            if(warehouse_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            let str = idrecord.toUpperCase();
            let newstr = str.replace(" ","");
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/warehousesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    warehouse_id: newstr,
                    warehouse_idbranch: warehouse_idbranch_field,
                    warehouse_name: warehouse_name_field,
                    warehouse_alamat: warehouse_alamat_field,
                    warehouse_phone1: warehouse_phone1_field,
                    warehouse_phone2: warehouse_phone2_field,
                    warehouse_status: warehouse_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }
                msgloader = json.message;
                RefreshHalaman()
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSave_storage() {
        let flag = true
        let msg = ""
        if(sDataStorage == "New"){
            if(storage_warehouse_id == ""){
                flag = false
                msg += "The Warehouse is required\n"
            }
            if(storage_id_field == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(storage_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(storage_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(storage_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(storage_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/warehousestoragesave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataStorage,
                    page:"CURR-SAVE",
                    warehousestorage_id: storage_id_field,
                    warehousestorage_idwarehouse: storage_warehouse_id,
                    warehousestorage_name: storage_name_field,
                    warehousestorage_status: storage_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataStorage=="New"){
                    clearField_storage()
                }
                msgloader = json.message;
                call_storage(storage_warehouse_id)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function handleSave_bin() {
        let flag = true
        let msg = ""
        if(sDataBin == "New"){
            if(bin_idstorage_field == ""){
                flag = false
                msg += "The Storage is required\n"
            }
            if(bin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(bin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(bin_name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(bin_status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/warehousestoragebinsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataBin,
                    page:"CURR-SAVE",
                    storagebin_id: parseInt(bin_id_field),
                    storagebin_idstorage: bin_idstorage_field,
                    storagebin_iduom: bin_iduom_field,
                    storagebin_name: bin_name_field,
                    storagebin_maxcapacity: parseFloat(bin_maxcapacity_field),
                    storagebin_status: bin_status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataBin=="New"){
                    clearField_bin()
                }
                msgloader = json.message;
                call_bin(bin_idstorage_field)
            } else if(json.status == 403){
                flag_btnsave = true;
                alert(json.message)
            } else {
                flag_btnsave = true;
                msgloader = json.message;
            }
            setTimeout(function () {
                css_loader = "display: none;";
            }, 1000);
        }else{
            alert(msg)
        }
    }
    async function call_storage(idwarehouse) {
        liststorage = [];
        total_liststorage = 0;
        const res = await fetch("/api/warehousestorage", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                warehouse_id: idwarehouse,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                let css_totalbin = ""
                total_liststorage = record.length;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(record[i]["warehousestorage_totalbin"] > 0){
                        css_totalbin = "color:blue;";
                    }else{
                        css_totalbin = "color:red;";
                    }
                    liststorage = [
                        ...liststorage,
                        {
                            warehousestorage_no: no,
                            warehousestorage_id: record[i]["warehousestorage_id"],
                            warehousestorage_name: record[i]["warehousestorage_name"],
                            warehousestorage_totalbin: record[i]["warehousestorage_totalbin"],
                            warehousestorage_totalbin_css: css_totalbin,
                            warehousestorage_status: record[i]["warehousestorage_status"],
                            warehousestorage_status_css: record[i]["warehousestorage_status_css"],
                            warehousestorage_create: record[i]["warehousestorage_create"],
                            warehousestorage_update: record[i]["warehousestorage_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_bin(idstorage) {
        listbin = [];
        total_listbin = 0;
        const res = await fetch("/api/warehousestoragebin", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                storage_id: idstorage,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                let css_maxcapacity = ""
                let css_totalcapacity = ""
                total_listbin = record.length;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    if(record[i]["warehousestoragebin_maxcapacity"] > 0){
                        css_maxcapacity = "color:blue;";
                    }else{
                        css_maxcapacity = "color:red;";
                    }
                    if(record[i]["warehousestoragebin_totalcapacity"] > 0){
                        css_totalcapacity = "color:blue;";
                    }else{
                        css_totalcapacity = "color:red;";
                    }
                    listbin = [
                        ...listbin,
                        {
                            warehousestoragebin_no: no,
                            warehousestoragebin_id: record[i]["warehousestoragebin_id"],
                            warehousestoragebin_iduom: record[i]["warehousestoragebin_iduom"],
                            warehousestoragebin_name: record[i]["warehousestoragebin_name"],
                            warehousestoragebin_maxcapacity: record[i]["warehousestoragebin_maxcapacity"],
                            warehousestoragebin_maxcapacity_css: css_maxcapacity,
                            warehousestoragebin_totalcapacity: record[i]["warehousestoragebin_totalcapacity"],
                            warehousestoragebin_totalcapacity_css: css_totalcapacity,
                            warehousestoragebin_status: record[i]["warehousestoragebin_status"],
                            warehousestoragebin_status_css: record[i]["warehousestoragebin_status_css"],
                            warehousestoragebin_create: record[i]["warehousestoragebin_create"],
                            warehousestoragebin_update: record[i]["warehousestoragebin_update"],
                        },
                    ];
                }
            }
        }
    }
    async function call_uom() {
        listuom = [];
        const res = await fetch("/api/uomshare", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listuom = [
                        ...listuom,
                        {
                            uom_id: record[i]["uom_id"],
                            uom_name: record[i]["uom_name"],
                        },
                    ];
                }
            }
        }
    }
    function clearField(){
        idrecord = "";
        flag_idbranch_field = false
        flag_id_field = false
        code_branch_display = ""
        warehouse_idbranch_field = "";
        warehouse_name_field = "";
        warehouse_alamat_field = "";
        warehouse_phone1_field = "";
        warehouse_phone2_field = "";
        warehouse_status_field = "";
        warehouse_create_field = "";
        warehouse_update_field = "";
    }
    function clearField_storage(){
        storage_flag_id = false;
        storage_id_field = "";
        storage_name_field = "";
        storage_status_field = "";
        storage_create_field = "";
        storage_update_field = "";
    }
    function clearField_bin(){
        bin_id_field = 0;
        bin_iduom_field = "";
        bin_name_field = "";
        bin_maxcapacity_field = 0;
        bin_status_field = "";
        bin_create_field = "";
        bin_update_field = "";

    }
    function callFunction(event){
        switch(event.detail){
            case "NEW":
                NewData("New","","","");
                break;
            case "REFRESH":
                RefreshHalaman();break;
            case "SAVE":
                handleSubmit();break;
        }
    }
    const handleChangeBranch = (e) => {
        code_branch_display = e.target.value
    };
    const handleChangeBranchTable = (e) => {
        alert(e.target.value)
    };
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterTafsirMimpi = [];
                listHome = [];
                const tafsir = {
                    searchTafsirMimpi,
                };
                dispatch("handleTafsirMimpi", tafsir);
        }  
    };
    
    
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
    }

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }

        if (searchStorage) {
            filterStorage = liststorage.filter(
                (item) =>
                    item.warehousestorage_name
                        .toLowerCase()
                        .includes(searchStorage.toLowerCase()) 
            );
        } else {
            filterStorage = [...liststorage];
        }
    }
</script>
<div id="loader" style="margin-left:50%;{css_loader}">
    {msgloader}
</div>
<div class="container" style="margin-top: 70px;">
    <div class="row">
        <div class="col-sm-12">
            <Button
                on:click={callFunction}
                button_function="NEW"
                button_title="<i class='bi bi-plus-lg'></i>&nbsp;New"
                button_css="btn-dark"/>
            <Button
                on:click={callFunction}
                button_function="REFRESH"
                button_title="<i class='bi bi-arrow-clockwise'></i>&nbsp;Refresh"
                button_css="btn-primary"/>
            <Panel
                card_search={true}
                card_title="{title_page}"
                card_footer={totalrecord}>
                <slot:template slot="card-search">
                    <div class="row" style="padding: 5px;">
                        <div class="col-lg-6">
                            <select
                                on:change="{handleChangeBranchTable}"
                                bind:value="{warehouse_idbranch_select}" 
                                class=" form-control ">
                                <option value="">--Please Select Branch--</option>
                                {#each listBranch as rec}
                                <option value="{rec.branch_id}">{rec.branch_name}</option>
                                {/each}
                            </select>
                        </div>
                        <div class="col-lg-6">
                            <input
                                bind:value={searchHome}
                                on:keypress={handleKeyboard_checkenter}
                                type="text"
                                class="form-control"
                                placeholder="Search Warehouse"
                                aria-label="Search"/>
                        </div>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BRANCH</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">WAREHOUSE</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                            //e,id,idbranch,name,alamat,phone1,phone2,status,create,update
                                                NewData("Edit",rec.home_id, rec.home_idbranch,
                                                rec.home_name,rec.home_alamat,rec.home_phone1,rec.home_phone2,
                                                rec.home_status, rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                            showStorage(rec.home_id);
                                            }} class="bi bi-box-seam"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmbranch}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                </tr>
                            {/each}
                        </tbody>
                        {:else}
                        <tbody>
                            <tr>
                                <td colspan="20">
                                    <center>
                                        <Loader />
                                    </center>
                                </td>
                            </tr>
                        </tbody>
                        {/if} 
                    </table>
                </slot:template>
            </Panel>
        </div>
    </div>
</div>

<Modal
	modal_id="modalentrycrud"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Branch</label>
                    <select
                        on:change="{handleChangeBranch}"
                        bind:value="{warehouse_idbranch_field}" 
                        name="currency" id="currency" 
                        disabled={flag_idbranch_field} 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listBranch as rec}
                        <option value="{rec.branch_id}">{rec.branch_name}</option>
                        {/each}
                    </select>
                    
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">CODE</label>
                    {#if flag_id_field == true}
                    <Input_custom
                        bind:value={idrecord}
                        input_tipe="text_uppercase_trim"
                        input_required="required"
                        input_maxlength="18"
                        disabled=true
                        input_placeholder="CODE"/>
                    {:else}
                    <div class="input-group mb-3">
                        <span class="input-group-text" id="basic-addon1">{code_branch_display}-</span>
                        <Input_custom
                            bind:value={idrecord}
                            input_tipe="text_uppercase_trim"
                            input_required="required"
                            input_maxlength="18"
                            input_placeholder="CODE"/>
                    </div>
                    
                    {/if}
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input_custom
                        bind:value={warehouse_name_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="60"
                        input_placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Alamat</label>
                    <textarea 
                        style="height: 100px;resize: none;" bind:value={warehouse_alamat_field} class="form-control required"/>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone1</label>
                    <Input_custom
                        bind:value={warehouse_phone1_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="20"
                        input_placeholder="Phone1"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Phone2</label>
                    <Input_custom
                        bind:value={warehouse_phone2_field}
                        input_tipe="text_standart"
                        input_maxlength="20"
                        input_placeholder="Phone2"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={warehouse_status_field}>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {warehouse_create_field}<br />
                        Update : {warehouse_update_field}
                    </div>
                </div>
                {/if}
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave();
            }} 
            button_function="SAVE"
            button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalliststorage"
	modal_size="modal-dialog-centered"
	modal_title="STORAGE | {storage_warehouse_title}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
    modal_search={true}
	modal_footer={true}>
    <slot:template slot="search">
        <input
            bind:value={searchStorage}
            type="text"
            class="form-control"
            placeholder="Search Storage"
            aria-label="Search"/>
    </slot:template>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan="2">&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CODE</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">STORAGE</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TOTAL BIN</th>
                </tr>
            </thead>
            {#if total_liststorage > 0}
                <tbody>
                    {#each filterStorage as rec }
                        <tr>
                            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                <i on:click={() => {
                                        //e,idstorage,nmstorage,statustorage
                                        call_formliststorage("Edit",rec.warehousestorage_id,rec.warehousestorage_name,rec.warehousestorage_status,
                                        rec.warehousestorage_create,rec.warehousestorage_update);
                                    }} class="bi bi-pencil"></i>
                            </td>
                            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                <i on:click={() => {
                                    showStorageBin(rec.warehousestorage_id);
                                    }} class="bi bi-box-seam"></i>
                            </td>
                            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.warehousestorage_no}</td>
                            <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.warehousestorage_status_css}">
                                    {status(rec.warehousestorage_status)}
                                </span>
                            </td>
                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.warehousestorage_id}</td>
                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.warehousestorage_name}</td>
                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.warehousestorage_totalbin_css}">{rec.warehousestorage_totalbin}</td>
                        </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="6">
                            <center>
                                <Loader />
                            </center>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                call_formliststorage("New");
            }} 
            button_title="<i class='bi bi-plus-lg'></i>&nbsp;New"
            button_css="btn-info"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalcrudstorage"
	modal_size="modal-dialog-centered"
	modal_title="Storage - {storage_warehouse_title+"  / "+sDataStorage}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">CODE</label>
            <div class="input-group mb-3">
                {#if sDataStorage == "New"}
                <span class="input-group-text" id="basic-addon1">{storage_warehouse_id}-</span>
                {/if}
                {#if storage_flag_id}
                    <Input_custom
                        bind:value={storage_id_field}
                        input_tipe="text_uppercase_trim"
                        input_required="required"
                        input_maxlength="10"
                        disabled=true
                        input_placeholder="CODE"/>
                {:else}
                    <Input_custom
                        bind:value={storage_id_field}
                        input_tipe="text_uppercase_trim"
                        input_required="required"
                        input_maxlength="10"
                        input_placeholder="CODE"/>
                {/if}
                
            </div>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Name</label>
            <Input_custom
                bind:value={storage_name_field}
                input_tipe="text_standart"
                input_required="required"
                input_maxlength="50"
                input_placeholder="Name"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Status</label>
            <select
                class="form-control required"
                bind:value={storage_status_field}>
                <option value="">--Please Select--</option>
                <option value="Y">ACTIVE</option>
                <option value="N">DEACTIVE</option>
            </select>
        </div>
        {#if sDataStorage != "New"}
            <div class="mb-3">
                <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                    Create : {storage_create_field}<br />
                    Update : {storage_update_field}
                </div>
            </div>
        {/if}
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_storage();
            }} 
            
            button_title="<i class='bi bi-save'></i>&nbsp;&nbspSave"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalliststoragebin"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="BIN | {bin_storage_id}"
    modal_body_css="height:500px; overflow-y: scroll;"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <table class="table table-striped ">
            <thead>
                <tr>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;">&nbsp;</th>
                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                    <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BIN</th>
                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UOM</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">MAX CAPACITY</th>
                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TOTAL CAPACITY</th>
                </tr>
            </thead>
            {#if total_listbin > 0}
                <tbody>
                    {#each listbin as rec }
                        <tr>
                            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                <i on:click={() => {
                                        //e,idstorage,nmstorage,statustorage
                                        // call_formlistbin = (e,idbin,iduom,nmbin,max,status,create,update)
                                        call_formlistbin("Edit",rec.warehousestoragebin_id,rec.warehousestoragebin_iduom,rec.warehousestoragebin_name,
                                        rec.warehousestoragebin_maxcapacity,rec.warehousestoragebin_status,
                                        rec.warehousestoragebin_create,rec.warehousestoragebin_update);
                                    }} class="bi bi-pencil"></i>
                            </td>
                            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.warehousestoragebin_no}</td>
                            <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.warehousestoragebin_status_css}">
                                    {status(rec.warehousestoragebin_status)}
                                </span>
                            </td>
                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.warehousestoragebin_name}</td>
                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.warehousestoragebin_iduom}</td>
                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.warehousestoragebin_maxcapacity_css}">{new Intl.NumberFormat().format(rec.warehousestoragebin_maxcapacity)}</td>
                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};{rec.warehousestoragebin_totalcapacity_css}">{new Intl.NumberFormat().format(rec.warehousestoragebin_totalcapacity)}</td>
                        </tr>
                    {/each}
                </tbody>
            {:else}
                <tbody>
                    <tr>
                        <td colspan="20">
                            <center>
                                <Loader />
                            </center>
                        </td>
                    </tr>
                </tbody>
            {/if}
        </table>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                call_formlistbin("New");
            }} 
            button_title="<i class='bi bi-plus-lg'></i>&nbsp;New"
            button_css="btn-info"/>
	</slot:template>
</Modal>

<Modal
	modal_id="modalcrudbin"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="Bin - {storage_warehouse_title+"  / "+sDataStorage}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input_custom
                        bind:value={bin_name_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={bin_status_field}>
                        <option value="">--Please Select--</option>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
            </div>
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Uom</label>
                    <select
                        bind:value="{bin_iduom_field}" 
                        class="form-control required">
                        <option value="">--Please Select--</option>
                        {#each listuom as rec}
                            <option value="{rec.uom_id}">{rec.uom_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Max Capacity</label>
                    <Input_custom
                        bind:value={bin_maxcapacity_field}
                        input_tipe="number_standart"
                        input_required="required"
                        input_maxlength="10"
                        input_placeholder="0"/>
                </div>
                {#if sDataBin != "New"}
                    <div class="mb-3">
                        <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                            Create : {bin_create_field}<br />
                            Update : {bin_update_field}
                        </div>
                    </div>
                {/if}
            </div>

        </div>
        
       
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_bin();
            }} 
            
            button_title="<i class='bi bi-save'></i>&nbsp;&nbspSave"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>