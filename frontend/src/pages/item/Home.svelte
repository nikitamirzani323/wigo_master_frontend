<script> 
    import Panel from "../../components/Panel.svelte";
    import Loader from "../../components/Loader.svelte";
	import Button from "../../components/Button.svelte";
    import Input_custom from '../../components/InputCustom.svelte' 
	import Modal from "../../components/Modal.svelte";
    import { createEventDispatcher } from "svelte";

    
	export let table_header_font = ""
	export let table_body_font = ""
	export let token = ""
	export let listHome = []
    export let listPage = [];
    export let listCateitem = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "ITEM"
    let sData = "";
    let myModal_newentry = "";
    let myModal_merek = "";
    let flag_id_field = false;
    let flag_btnsave = true;
    let listMerek = [];
    let idmerek_field = 0;
    let nmmerek_field = "";
    let idcateitem_field = "";
    let iduom_field = "";
    let name_field = "";
    let descp_field = "";
    let urlimg_field = "";
    let inventory_field = false;
    let sales_field = false;
    let purchase_field = false;
    let status_field = "";
    let create_field = "";
    let update_field = "";
    let idrecord = 0;
    let pagingnow = 0;
    let searchMerek = "";
    let searchHome = "";
    let filterMerek = [];
    let filterHome = [];
    let listuom = [];
    
    //==ITEMUOM===
    let sDataItemUom = "";
    let list_itemuom = [];
    let total_itemuom = 0;
    let id_itemuom = 0;
    let iduom_itemuom = "";
    let iduom_flag_itemuom = false;
    let iduomdefault_itemuom = "";
    let iditem_itemuom = "";
    let default_itemuom = "N";
    let convertion_itemuom = 1;
    let create_itemuom = "";
    let update_itemuom = "";


    let css_loader = "display: none;";
    let msgloader = "";

    $: {
        if (searchHome) {
            filterHome = listHome.filter(
                (item) =>
                    item.home_name
                        .toLowerCase()
                        .includes(searchHome.toLowerCase()) || 
                    item.home_id
                        .toLowerCase()
                        .includes(searchHome.toLowerCase())
            );
        } else {
            filterHome = [...listHome];
        }
        if (searchMerek) {
            filterMerek = listMerek.filter(
                (item) =>
                    item.merek_name
                        .toLowerCase()
                        .includes(searchMerek.toLowerCase()) 
            );
        } else {
            filterMerek = [...listMerek];
        }
    }
    
    const NewData = (e,id,idmerek,nmmerek,idcateitem,nmitem,descp,urlimg,status,purchase,sales,inventory,create,update) => {
        sData = e
        if(sData == "New"){
            call_uom()
            clearField()
            myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
            myModal_newentry.show();
        }else{
            call_itemuom(id)
            flag_id_field = true;
            idrecord = id
            idmerek_field = idmerek;
            nmmerek_field = nmmerek;
            idcateitem_field = idcateitem;
            name_field = nmitem;
            descp_field = descp;
            urlimg_field = urlimg;
            inventory_field = inventory=="Y"?true:false;
            sales_field = sales=="Y"?true:false;
            purchase_field = purchase=="Y"?true:false;
            status_field = status;
            create_field = create;
            update_field = update;
            myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrudedit"));
            myModal_newentry.show();
        }
    };
    const NewItemUom = (e,id,iduom,default_uom,convertion,create,update) => {
        sDataItemUom = e
        call_uom()
        if(sDataItemUom == "Edit"){
            iduom_flag_itemuom = true;
            id_itemuom = id;
            iduom_itemuom = iduom;
            default_itemuom = default_uom;
            convertion_itemuom = parseFloat(convertion);
            create_itemuom = create;
            update_itemuom = update;
        }else{
            clearField_iteuom()
        }
        


        myModal_newentry = new bootstrap.Modal(document.getElementById("modalcruditemuom"));
        myModal_newentry.show();
        
    };
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        
        if(sData == "New"){
            if(parseInt(idmerek_field) < 0){
                flag = false
                msg += "The Merek is required\n"
            }
            if(idcateitem_field == ""){
                flag = false
                msg += "The Category is required\n"
            }
            if(iduom_field == ""){
                flag = false
                msg += "The UOM is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }else{
            if(parseInt(idmerek_field) < 0){
                flag = false
                msg += "The Merek is required\n"
            }
            if(idcateitem_field == ""){
                flag = false
                msg += "The Category is required\n"
            }
            if(name_field == ""){
                flag = false
                msg += "The Name is required\n"
            }
            if(status_field == ""){
                flag = false
                msg += "The Status is required\n"
            }
        }
        
        if(flag){
            if(inventory_field){
                inventory_field = "Y"
            }else{
                inventory_field = "N"
            }
            if(sales_field){
                sales_field = "Y"
            }else{
                sales_field = "N"
            }
            if(purchase_field){
                purchase_field = "Y"
            }else{
                purchase_field = "N"
            }
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/itemsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    item_search: searchHome,
                    item_page: parseInt(pagingnow),
                    item_id: idrecord,
                    item_idmerek: parseInt(idmerek_field),
                    item_idcateitem: idcateitem_field,
                    item_iduom: iduom_field,
                    item_name: name_field,
                    item_descp: descp_field,
                    item_urlimg: urlimg_field,
                    item_inventory: inventory_field,
                    item_sales: sales_field,
                    item_purchase: purchase_field,
                    item_status: status_field,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sData=="New"){
                    clearField()
                }else{
                    if(inventory_field == "Y"){
                        inventory_field = true
                    }else{
                        inventory_field = false
                    }
                    if(sales_field== "Y"){
                        sales_field = true
                    }else{
                        sales_field = false
                    }
                    if(purchase_field== "Y"){
                        purchase_field = true
                    }else{
                        purchase_field = false
                    }
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
    async function handleSave_itemuom() {
        let flag = true
        let msg = ""
        
        if(sDataItemUom == "New"){
            if(iditem_itemuom == ""){
                flag = false
                msg += "The Item is required\n"
            }
            if(iduom_itemuom == ""){
                flag = false
                msg += "The UOM is required\n"
            }
           
            if(parseFloat(convertion_itemuom) <= 0){
                flag = false
                msg += "The Convertion must be greater than 0\n"
            }
        }else{
            if(iditem_itemuom == ""){
                flag = false
                msg += "The Item is required\n"
            }
            if(iduom_itemuom == ""){
                flag = false
                msg += "The UOM is required\n"
            }
           
            if(parseFloat(convertion_itemuom) <= 0){
                flag = false
                msg += "The Convertion must be greater than 0\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/itemuomsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sDataItemUom,
                    page:"CURR-SAVE",
                    itemuom_search: searchHome,
                    itemuom_page: parseInt(pagingnow),
                    itemuom_id: parseInt(id_itemuom),
                    itemuom_iditem: iditem_itemuom,
                    itemuom_iduom: iduom_itemuom,
                    itemuom_default: default_itemuom,
                    itemuom_conversion: parseFloat(convertion_itemuom),
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                if(sDataItemUom=="New"){
                    clearField_iteuom()
                }
                call_itemuom(iditem_itemuom)
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
    async function handleDelete_itemuom(iditemuom) {
        let flag = true
        let msg = ""
        
        if(iditemuom == ""){
            flag = false
            msg += "The Item is required\n"
        }
        if(iditem_itemuom == ""){
            flag = false
            msg += "The item is required\n"
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            const res = await fetch("/api/itemuomdelete", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    page:"CURR-SAVE",
                    itemuom_search: searchHome,
                    itemuom_page: parseInt(pagingnow),
                    itemuom_id: parseInt(iditemuom),
                    itemuom_iditem: iditem_itemuom,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                call_itemuom(iditem_itemuom)
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
    const ShowMerek = () => {
        call_merek()
        myModal_merek = new bootstrap.Modal(document.getElementById("modalmerek"));
        myModal_merek.show();
    };
    const handle_pilihmerek = (e,nm) => {    
        idmerek_field = e;
        nmmerek_field = nm;
        myModal_merek.hide();
    };
    async function call_merek() {
        listMerek = [];
        const res = await fetch("/api/merekshare", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                merek_search: searchMerek,
                merek_page : 0
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    listMerek = [
                        ...listMerek,
                        {
                            merek_id: record[i]["merek_id"],
                            merek_name: record[i]["merek_name"],
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
    async function call_itemuom(iditem) {
        list_itemuom = [];
        iditem_itemuom = iditem;
        const res = await fetch("/api/itemuom", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                item_id: iditem,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                let tipe = ""
                total_itemuom = record.length;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    
                    if(record[i]["itemuom_default"] == "Y"){
                        iduomdefault_itemuom = record[i]["itemuom_iduom"]
                        tipe = "Selling"
                    }else{
                        tipe = "Purchasing"
                    }
                    list_itemuom = [
                        ...list_itemuom,
                        {
                            itemuom_no: no,
                            itemuom_id: record[i]["itemuom_id"],
                            itemuom_iduom: record[i]["itemuom_iduom"],
                            itemuom_nmuom: record[i]["itemuom_nmuom"],
                            itemuom_tipe: tipe,
                            itemuom_default: record[i]["itemuom_default"],
                            itemuom_default_css: record[i]["itemuom_default_css"],
                            itemuom_conversion: record[i]["itemuom_conversion"],
                            itemuom_create: record[i]["itemuom_create"],
                            itemuom_update: record[i]["itemuom_update"],
                        },
                    ];
                }
            }
        }
    }
    function clearField(){
        idrecord = "";
        idmerek_field = 0;
        nmmerek_field = "";
        idcateitem_field = "";
        iduom_field = "";
        name_field = "";
        descp_field = "";
        urlimg_field = "";
        inventory_field = false;
        sales_field = false;
        purchase_field = false;
        status_field = "";
        create_field = "";
        update_field = "";
    }
    function clearField_iteuom(){
        id_itemuom = 0;
        iduom_flag_itemuom = false;
        iduom_itemuom = "";
        default_itemuom = "N";
        convertion_itemuom = 1;
        create_itemuom = "";
        update_itemuom = "";
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
    const handleKeyboard_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
                filterHome = [];
                listHome = [];
                const searchdata = {
                  searchHome,
                };
                dispatch("handleSearch", searchdata);
        }  
    };
    const handleKeyboard_merek_checkenter = (e) => {
        let keyCode = e.which || e.keyCode;
        if (keyCode === 13) {
            filterMerek = [];
            listMerek = [];
            call_merek();
        }
    };
    const handleSelectPaging = (event) => {
      let page = event.target.value;
      pagingnow = page;
      const pattern = {
        page
      };
      dispatch("handlePaging", pattern);
    };
    function status(e){
        let result = "DEACTIVE"
        if(e == "Y"){
            result = "ACTIVE"
        }
        return result
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
                <slot:template slot="card-title">
                    <div class="float-end">
                      <select
                        on:change={handleSelectPaging}
                        style="text-align: center;"
                        class="form-control">
                        {#each listPage as rec}
                          <option value={rec.page_value}>{rec.page_display}</option>
                        {/each}
                      </select>
                    </div>
                  </slot:template>
                <slot:template slot="card-search">
                    <div class="col-lg-12" style="padding: 5px;">
                        <input
                            bind:value={searchHome}
                            on:keypress={handleKeyboard_checkenter}
                            type="text"
                            class="form-control"
                            placeholder="Search Code, Item"
                            aria-label="Search"/>
                    </div>
                </slot:template>
                <slot:template slot="card-body">
                    <table class="table table-striped ">
                        <thead>
                            <tr>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" >&nbsp;</th>
                                <th NOWRAP width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
                                <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">&nbsp;</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CODE</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CATEGORY ITEM</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">MEREK</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">ITEM</th>
                                <th NOWRAP width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEFAULT UOM</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">PURCHASE</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">SALES</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">INVENTORY</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                            //e,id,idcateitem,nmitem,descp,status,purchase,sales,inventory,create,update
                                                NewData("Edit",rec.home_id, rec.home_idmerek,rec.home_nmmerek ,rec.home_idcateitem,
                                                rec.home_name,rec.home_descp,rec.home_urlimg,rec.home_status,
                                                rec.home_purchase,rec.home_sales,rec.home_inventory,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {status(rec.home_status)}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmcateitem}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmmerek}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_name}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_iduom}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_purchase_css}">
                                            {rec.home_purchase}
                                        </span>
                                    </td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_sales_css}">
                                            {rec.home_sales}
                                        </span>
                                    </td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_inventory_css}">
                                            {rec.home_inventory}
                                        </span>
                                    </td>
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
                    <label for="exampleForm" class="form-label">Category Item</label>
                    <select
                        bind:value="{idcateitem_field}" 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listCateitem as rec}
                        <option value="{rec.cateitem_id}">{rec.cateitem_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Merek</label>
                    <div class="input-group mb-3">
                        <input type="text" 
                            bind:value="{nmmerek_field}" 
                            disabled
                            class="form-control" placeholder="Merek" >
                        <Button on:click={() => {
                                ShowMerek();
                            }} 
                            button_function="New"
                            button_title="<i class='bi bi-search'></i>"
                            button_css="btn-warning"/>
                    </div>
                </div>
                {#if sData == "New"}
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Default Uom - Selling</label>
                    <select
                        bind:value="{iduom_field}" 
                        class="form-control required">
                        <option value="">--Please Select--</option>
                        {#each listuom as rec}
                            <option value="{rec.uom_id}">{rec.uom_name}</option>
                        {/each}
                    </select>
                </div>
                {/if}
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input_custom
                        bind:value={name_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Description</label>
                    <textarea 
                        style="height: 100px;resize: none;" bind:value={descp_field} class="form-control "/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url Image</label>
                    <Input_custom
                        bind:value={urlimg_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="500"
                        input_placeholder="Url Image"/>
                </div>
                <div class="mb-3">
                    <img width="100" src="{urlimg_field}" alt="">
                </div>
            </div>
            <div class="col-md-6">
                <div class="form-check">
                    <input 
                        bind:checked={purchase_field}
                        class="form-check-input" 
                        type="checkbox" value="" id="flexCheckChecked" >
                    <label class="form-check-label" for="flexCheckDefault">
                        Purchase
                    </label>
                </div>
                <div class="form-check">
                    <input 
                        bind:checked={sales_field}
                        class="form-check-input" 
                        type="checkbox" value="" id="flexCheckChecked" >
                    <label class="form-check-label" for="flexCheckChecked">
                      Sales
                    </label>
                </div>
                <div class="form-check">
                    <input 
                        bind:checked={inventory_field}
                        class="form-check-input" 
                        type="checkbox" value="" id="flexCheckChecked" >
                    <label class="form-check-label" for="flexCheckChecked">
                      Inventory
                    </label>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="">--Please Select--</option>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {create_field}<br />
                        Update : {update_field}
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
	modal_id="modalentrycrudedit"
	modal_size="modal-dialog-centered modal-lg"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={false}>
	<slot:template slot="body">
        <div class="row">
            <div class="col-md-6">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Category Item</label>
                    <select
                        bind:value="{idcateitem_field}" 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listCateitem as rec}
                        <option value="{rec.cateitem_id}">{rec.cateitem_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Merek</label>
                    <div class="input-group mb-3">
                        <input type="text" 
                            bind:value="{nmmerek_field}" 
                            disabled
                            class="form-control" placeholder="Merek" >
                        <Button on:click={() => {
                                ShowMerek();
                            }} 
                            button_function="New"
                            button_title="<i class='bi bi-search'></i>"
                            button_css="btn-warning"/>
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Name</label>
                    <Input_custom
                        bind:value={name_field}
                        input_tipe="text_standart"
                        input_required="required"
                        input_maxlength="50"
                        input_placeholder="Name"/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Description</label>
                    <textarea 
                        style="height: 100px;resize: none;" bind:value={descp_field} class="form-control "/>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Url Image</label>
                    <Input_custom
                        bind:value={urlimg_field}
                        input_tipe="text_standart"
                        input_required=""
                        input_maxlength="500"
                        input_placeholder="Url Image"/>
                </div>
                <div class="mb-3">
                    <img width="150" src="{urlimg_field}" alt="">
                </div>
            </div>
            <div class="col-md-6">
                <div class="form-check">
                    <input 
                        bind:checked={purchase_field}
                        class="form-check-input" 
                        type="checkbox" value="" id="flexCheckChecked" >
                    <label class="form-check-label" for="flexCheckDefault">
                        Purchase
                    </label>
                </div>
                <div class="form-check">
                    <input 
                        bind:checked={sales_field}
                        class="form-check-input" 
                        type="checkbox" value="" id="flexCheckChecked" >
                    <label class="form-check-label" for="flexCheckChecked">
                      Sales
                    </label>
                </div>
                <div class="form-check">
                    <input 
                        bind:checked={inventory_field}
                        class="form-check-input" 
                        type="checkbox" value="" id="flexCheckChecked" >
                    <label class="form-check-label" for="flexCheckChecked">
                      Inventory
                    </label>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Status</label>
                    <select
                        class="form-control required"
                        bind:value={status_field}>
                        <option value="">--Please Select--</option>
                        <option value="Y">ACTIVE</option>
                        <option value="N">DEACTIVE</option>
                    </select>
                </div>
                {#if sData != "New"}
                <div class="mb-3">
                    <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                        Create : {create_field}<br />
                        Update : {update_field}
                    </div>
                </div>
                {/if}
                <div class="mb-3">
                    <div class="float-end">
                        {#if flag_btnsave}
                            <Button on:click={() => {
                                    handleSave();
                                }} 
                                button_function="SAVE"
                                button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
                                button_css="btn-warning"/>
                        {/if}
                    </div>
                </div>
            </div>
        </div>
        <div class="accordion" id="accordionExample">
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseOne" aria-expanded="true" aria-controls="collapseOne">
                        Setting Uom and conversions
                    </button>
                </h2>
                <div id="collapseOne" class="accordion-collapse collapse show" data-bs-parent="#accordionExample">
                    <div class="accordion-body">
                        {#if parseInt(total_itemuom) < 3}
                        <div class="float-end">
                            <Button on:click={() => {
                                    NewItemUom("New",iditem_itemuom);
                                }} 
                                button_function=""
                                button_title="<i class='bi bi-plus-lg'></i>&nbsp;New Uom"
                                button_css="btn-dark"/>
                        </div>
                        {/if}
                        <table class="table table-striped ">
                            <thead>
                                <tr>
                                    <th NOWRAP width="1%" style="text-align: center;vertical-align: top;" colspan=2>&nbsp;</th>
                                    <th NOWRAP width="2%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEFAULT</th>
                                    <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                    <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UOM</th>
                                    <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CONVERTION</th>
                                </tr>
                            </thead>
                            {#if total_itemuom > 0}
                                <tbody>
                                    {#each list_itemuom as rec }
                                        <tr>
                                            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                                <i on:click={() => {
                                                        //e,idstorage,nmstorage,statustorage
                                                        handleDelete_itemuom(rec.itemuom_id);
                                                    }} class="bi bi-trash"></i>
                                            </td>
                                            <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                                <i on:click={() => {
                                                        //e,idstorage,nmstorage,statustorage
                                                        NewItemUom("Edit",rec.itemuom_id,rec.itemuom_iduom,rec.itemuom_default,rec.itemuom_conversion,
                                                        rec.itemuom_create,rec.itemuom_update);
                                                    }} class="bi bi-pencil"></i>
                                            </td>
                                            <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                                <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.itemuom_default_css}">
                                                    {rec.itemuom_default}
                                                </span>
                                            </td>
                                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.itemuom_tipe}</td>
                                            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.itemuom_nmuom}</td>
                                            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{rec.itemuom_conversion} {iduomdefault_itemuom}</td>
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
                    </div>
                </div>
            </div>
            <div class="accordion-item">
                <h2 class="accordion-header">
                    <button class="accordion-button" type="button" data-bs-toggle="collapse" data-bs-target="#collapseTwo" aria-expanded="false" aria-controls="collapseTwo">
                        Inventory
                    </button>
                </h2>
                <div id="collapseTwo" class="accordion-collapse collapse" data-bs-parent="#accordionExample">
                    <div class="accordion-body">
                        
                    </div>
                </div>
            </div>
        </div>
        
	</slot:template>
	<slot:template slot="footer">
        
	</slot:template>
</Modal>

<Modal
	modal_id="modalcruditemuom"
	modal_size="modal-dialog-centered"
	modal_title="Uom - {iditem_itemuom}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Uom</label>
            <select
                bind:value="{iduom_itemuom}" 
                disabled={iduom_flag_itemuom}
                class="form-control required">
                <option value="">--Please Select--</option>
                {#each listuom as rec}
                    <option value="{rec.uom_id}">{rec.uom_name}</option>
                {/each}
            </select>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Convertion</label>
            <Input_custom
                bind:value={convertion_itemuom}
                input_tipe="number_standart"
                input_required="required"
                input_maxlength="10"
                input_placeholder="0"/>
        </div>
        <div class="mb-3">
            <label for="exampleForm" class="form-label">Default Uom</label>
            <select
                class="form-control required"
                bind:value={default_itemuom}>
                <option value="N">N</option>
                <option value="Y">Y</option>
            </select>
        </div>
        {#if sDataItemUom != "New"}
        <div class="mb-3">
            <div class="alert alert-secondary" style="font-size: 11px; padding:10px;" role="alert">
                Create : {create_itemuom}<br />
                Update : {update_itemuom}
            </div>
        </div>
        {/if}
       
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleSave_itemuom();
            }} 
            
            button_title="<i class='bi bi-save'></i>&nbsp;&nbspSave"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
  modal_id="modalmerek"
  modal_size="modal-dialog-centered"
  modal_title="MEREK"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_search={true}
  modal_footer={false}>
    <slot:template slot="search">
        <input
            bind:value={searchMerek}
            on:keypress={handleKeyboard_merek_checkenter}
            type="text"
            class="form-control"
            placeholder="Search Merek"
            aria-label="Search"/>
    </slot:template>
    <slot:template slot="body">
        <table class="table table-sm">
            <thead>
                <tr>
                    <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">MEREK</th>
                </tr>
            </thead>
            <tbody>
                {#each filterMerek as rec}
                    <tr style="cursor: pointer;" on:click={() => {
                        handle_pilihmerek(rec.merek_id,rec.merek_name);
                        }} >
                        <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};text-decoration:underline;">{rec.merek_name}</td>
                    </tr>
                {/each}
            </tbody>
        </table>
    </slot:template>
</Modal>