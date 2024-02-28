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
    export let listCurr = [];
    export let listBranch = [];
    export let listDepartement = [];
    export let listPage = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "PURCHASE REQUEST"
    let sData = "";
    let myModal_newentry = "";
    let myModal_item = "";
    let myModal_employee = "";
    let flag_id_field = false;
    let flag_btnsave = true;
    let lock_document = false;
    let iddepartement_field = "";
    let idemployee_field = "";
    let nmemployee_field = "";
    let idbranch_field = "";
    let idcurr_field = "";
    let tipedoc_field = "";
    let listdetail_field = [];
    let totalitem_field = 0;
    let subtotal_field = 0;
    let remark_field = "";
    let status_field = "";
    let create_field = "";
    let update_field = "";

    //===EMPLOYEE==
    let listEmployee = []

    //===ITEM==
    let listitem = []
    let listuom = []
    let iditem_item_field = ""
    let nmitemdisplay_item_field = ""
    let nmitem_item_field = ""
    let desc_item_field = ""
    let qty_item_field = 0
    let iduom_item_field = ""
    let price_item_field = 0
    let total_item_field = 0
    let purpose_item_field = ""
    let subtotal_detail = 0

    let idrecord = "";
    let pagingnow = 0;
    let searchHome = "";
    let filterHome = [];
    let css_loader = "display: none;";
    let msgloader = "";

    let page_left = "col-md-4"
    let page_right = "col-md-8"
    let page_toggle = true;

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
    }
    const pageFullScreen = () =>{
        
        if(page_toggle){
            page_left = "d-none"
            page_right = "col-md-12"
            page_toggle = false
        }else{
            page_left = "col-md-4"
            page_right = "col-md-8"
            page_toggle = true
        }   
        
    }
    const NewData = (e,id,iddepartement,idemployee,nmemployee,idbranch,curr,tipedoc,remark,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
            lock_document = true;
        }else{
            call_detail(id)
            lock_document = false;
            flag_id_field = true;
            idrecord = id
            iddepartement_field = iddepartement;
            idemployee_field = idemployee;
            nmemployee_field = idemployee+"-"+nmemployee;
            idbranch_field = idbranch;
            idcurr_field = curr;
            tipedoc_field = tipedoc;
            remark_field = remark;
            status_field = status;
            create_field = create;
            update_field = update;
            if(status == "OPEN"){
                lock_document = true;
            }
        }
        myModal_newentry = new bootstrap.Modal(document.getElementById("modalentrycrud"));
        myModal_newentry.show();
    };
   
    const RefreshHalaman = () => {
        dispatch("handleRefreshData", "call");
    };
    async function handleSave() {
        let flag = true
        let msg = ""
        
        if(sData == "New"){
            if(idbranch_field == ""){
                flag = false
                msg += "The Branch is required\n"
            }
            if(iddepartement_field == ""){
                flag = false
                msg += "The Departement is required\n"
            }
            if(idemployee_field == ""){
                flag = false
                msg += "The Employee is required\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Currency is required\n"
            }
            if(tipedoc_field == ""){
                flag = false
                msg += "The Document is required\n"
            }
            if(listdetail_field == ""){
                flag = false
                msg += "The Detail is required\n"
            }
        }else{
            if(idrecord == ""){
                flag = false
                msg += "The Code is required\n"
            }
            if(idbranch_field == ""){
                flag = false
                msg += "The Branch is required\n"
            }
            if(iddepartement_field == ""){
                flag = false
                msg += "The Departement is required\n"
            }
            if(idemployee_field == ""){
                flag = false
                msg += "The Employee is required\n"
            }
            if(idcurr_field == ""){
                flag = false
                msg += "The Currency is required\n"
            }
            if(tipedoc_field == ""){
                flag = false
                msg += "The Document is required\n"
            }
            if(listdetail_field == ""){
                flag = false
                msg += "The Detail is required\n"
            }
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            totalitem_field = listdetail_field.length
            const res = await fetch("/api/purchaserequestsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    purchaserequest_search: searchHome,
                    purchaserequest_page: parseInt(pagingnow),
                    purchaserequest_id: idrecord,
                    purchaserequest_idbranch: idbranch_field,
                    purchaserequest_iddepartement: iddepartement_field,
                    purchaserequest_idemployee: idemployee_field,
                    purchaserequest_idcurr: idcurr_field,
                    purchaserequest_tipedoc: tipedoc_field,
                    purchaserequest_listdetail: listdetail_field,
                    purchaserequest_totalitem: parseFloat(totalitem_field),
                    purchaserequest_subtotal: parseFloat(subtotal_detail),
                    purchaserequest_remark: remark_field,
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
    async function handleStatusPurchaseRequest(e) {
        let flag = true
        let msg = ""
        if(idrecord == ""){
            flag = false
            msg += "The Document is required\n"
        }
        if(e == ""){
            flag = false
            msg += "The Branch is required\n"
        }
       
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            totalitem_field = listdetail_field.length
            const res = await fetch("/api/purchaserequeststatussave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    purchaserequest_id: idrecord,
                    purchaserequest_status: e,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
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
    function clearField(){
        idrecord = "";
        lock_document = false;
        iddepartement_field = "";
        idemployee_field = "";
        nmemployee_field = "";
        idbranch_field = "";
        idcurr_field = "";
        tipedoc_field = "";
        listdetail_field = [];
        remark_field = "";
        subtotal_detail = 0;
        status_field = "";
        create_field = "";
        update_field = "";
    }
    const handleChangeDepartement = (e) => {
        // alert(e.target.value)
        idemployee_field = "";
        nmemployee_field = "";
    };
    const ShowEmployee = () => {
        if(iddepartement_field == ""){
            alert("the Departement is required")
        }else{
            call_employee(iddepartement_field)
            myModal_employee = new bootstrap.Modal(document.getElementById("modalemployee"));
            myModal_employee.show();
        }
    };
    const handle_pilihemployee = (e,nm) => {
        idemployee_field = e
        nmemployee_field = e+" - "+nm
        
        myModal_employee.hide();
    };
    const ShowItem = () => {
        iditem_item_field = ""
        nmitem_item_field = ""
        iduom_item_field = ""
        listuom = []
        purpose_item_field = ""
        qty_item_field = 0
        price_item_field = 0
        call_item("")
        myModal_item = new bootstrap.Modal(document.getElementById("modalitem"));
        myModal_item.show();
    };
    const handle_pilihitem = (e,nm,uom) => {
        iditem_item_field = e
        nmitem_item_field = nm
        nmitemdisplay_item_field = e+" - "+nm
        iduom_item_field = ""
        listuom = []
        for(let i=0;i<uom.length;i++){
            listuom = [
                    ...listuom,
                    {
                        itemuom_iduom: uom[i]["itemuom_iduom"],
                    },
                ];
        }

     
        myModal_item.hide();
    };
    const handleListDetail = () => {
        let flag = true;
        let msg = "";
        if(iditem_item_field == "" || nmitem_item_field == ""){
            flag = false
            msg += "The Item is required\n"
        }
        if(iduom_item_field == ""){
            flag = false
            msg += "The Uom is required\n"
        }
        if(parseFloat(qty_item_field)<1){
            flag = false
            msg += "The Qty is required\n"
        }
        if(parseFloat(price_item_field)<1){
            flag = false
            msg += "The Price is required\n"
        }

        for(let i=0;i<listdetail_field.length;i++){
            if(listdetail_field[i].detail_iditem == iditem_item_field){
                flag = false
                msg += "Duplicate item\n"
                break;
            }
        }
    

        if(flag){
            let total = parseInt(qty_item_field)* parseInt(price_item_field);
            subtotal_detail = subtotal_detail + total;
            listdetail_field = [
                ...listdetail_field,
                {
                    detail_iditem: iditem_item_field,
                    detail_nmitem: nmitem_item_field,
                    detail_qty: parseFloat(qty_item_field),
                    detail_iduom: iduom_item_field,
                    detail_price: parseFloat(price_item_field),
                    detail_total: parseFloat(total),
                    detail_purpose: purpose_item_field,
                },
            ];

            listuom = []
            iditem_item_field = ""
            nmitemdisplay_item_field = ""
            nmitem_item_field = ""
            desc_item_field = ""
            qty_item_field = 0
            iduom_item_field = ""
            price_item_field = 0
            purpose_item_field = ""
        }else{
            alert(msg)
        }
    }
    const handleDeleteListDetail = (e) => {
        let temp = listdetail_field.filter(
            (item) => item.detail_iditem !== e
        );
        listdetail_field = [];
        subtotal_detail = 0;
        for (var i = 0; i < temp.length; i++) {
            let total = parseFloat(temp[i].detail_qty)* parseFloat(temp[i].detail_price);
            subtotal_detail = subtotal_detail + total;

            listdetail_field = [
            ...listdetail_field,
                {
                    detail_iditem: temp[i].detail_iditem,
                    detail_nmitem: temp[i].detail_nmitem,
                    detail_qty: temp[i].detail_qty,
                    detail_iduom: temp[i].detail_iduom,
                    detail_price: temp[i].detail_price,
                    detail_total: total,
                    detail_purpose: temp[i].detail_purpose,
                },
            ];
        }
    }
    const ShowFormDetail = () => {
        if(tipedoc_field == ""){
            alert("the Tipe Document is required")
        }else{
            if(tipedoc_field == "ITEM"){
                myModal_newentry = new bootstrap.Modal(document.getElementById("modalcruddetail"));
                myModal_newentry.show();
            }else{
                myModal_newentry = new bootstrap.Modal(document.getElementById("modalcruddetailservice"));
                myModal_newentry.show();
            }
            
        }
    };
    async function call_employee(iddepart) {
        listEmployee = [];
        const res = await fetch("/api/employeeshare", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
            employee_iddepartement: iddepart,
        }),
        });
        const json = await res.json();
        if (json.status == 200) {
        let record = json.record;
        if (record != null) {
            let totalemployee = record.length;
            let no = 0;
            for (var i = 0; i < record.length; i++) {
            no = no + 1;
            listEmployee = [
                ...listEmployee,
                {
                    employee_no: no,
                    employee_id: record[i]["employee_id"],
                    employee_name: record[i]["employee_name"],
                },
            ];
            }
        }
        }
    }
    async function call_item(searchitem) {
        listitem = [];
        const res = await fetch("/api/itemshare", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                item_search: searchitem,
                item_page: 0,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                no = no + 1;
                listitem = [
                    ...listitem,
                    {
                        itemshare_no: no,
                        itemshare_id: record[i]["itemshare_id"],
                        itemshare_nmcateitem: record[i]["itemshare_nmcateitem"],
                        itemshare_name: record[i]["itemshare_name"],
                        itemshare_descp: record[i]["itemshare_descp"],
                        itemshare_urlimg: record[i]["itemshare_urlimg"],
                        itemshare_uom: record[i]["itemshare_uom"],
                    },
                ];
                }
            }
        }
    }
    async function call_detail(idpurchase) {
        listdetail_field = [];
        subtotal_detail = 0;
        const res = await fetch("/api/purchaserequestdetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                purchaserequest_id: idpurchase,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    let total = parseFloat(record[i]["purchaserequestdetail_qty"]) * parseFloat(record[i]["purchaserequestdetail_price"])
                    subtotal_detail = subtotal_detail + total;
                    listdetail_field = [
                        ...listdetail_field,
                        {
                            detail_no: no,
                            detail_iditem: record[i]["purchaserequestdetail_iditem"],
                            detail_nmitem: record[i]["purchaserequestdetail_nmitem"],
                            detail_purpose: record[i]["purchaserequestdetail_purpose"],
                            detail_qty: record[i]["purchaserequestdetail_qty"],
                            detail_iduom: record[i]["purchaserequestdetail_iduom"],
                            detail_price: record[i]["purchaserequestdetail_price"],
                            detail_total: total,
                        },
                    ];
                }
            }
        }
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
<div class="container-fluid" style="margin-top: 70px;">
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
                            placeholder="Search Code"
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
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DOCUMENT</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DATE</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TIPE</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">BRANCH</th>
                                <th NOWRAP width="5%" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">DEPARTEMENT</th>
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">EMPLOYEE</th>
                                <th NOWRAP width="5%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">TOTAL ITEM</th>
                                <th NOWRAP width="5%" style="text-align: center;vertical-align: top;font-weight:bold;font-size: {table_header_font};">CURR</th>
                                <th NOWRAP width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size: {table_header_font};">SUBTOTAL</th>
                            </tr>
                        </thead>
                        {#if totalrecord > 0}
                        <tbody>
                            {#each filterHome as rec }
                                <tr>
                                    <td NOWRAP style="text-align: center;vertical-align: top;cursor:pointer;">
                                        <i on:click={() => {
                                            //e,id,iddepartement,idemployee,nmemployee,idbranch,curr,tipedoc,remark,status,create,update
                                                NewData("Edit",rec.home_id, rec.home_iddepartement,rec.home_idemployee,rec.home_nmemployee,
                                                rec.home_idbranch,rec.home_idcurr,rec.home_tipedoc,rec.home_remark,rec.home_status,
                                                rec.home_create, rec.home_update);
                                            }} class="bi bi-pencil"></i>
                                    </td>
                                    <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_no}</td>
                                    <td NOWRAP  style="text-align: center;vertical-align: top;font-size: 11px;">
                                        <span style="padding: 5px;border-radius: 10px;padding-right:10px;padding-left:10px;{rec.home_status_css}">
                                            {rec.home_status}
                                        </span>
                                    </td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_id}</td>
                                    <td  NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_date}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_tipedoc}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmbranch}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmdepartement}</td>
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmemployee}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_totalitem)}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_idcurr}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_totalpr)}</td>
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
	modal_size="modal-dialog-centered modal-xl"
	modal_title="{title_page+"/"+sData}"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="{page_left} ">
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Departement</label>
                    <select
                        on:change="{handleChangeDepartement}"
                        bind:value="{iddepartement_field}" 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listDepartement as rec}
                        <option value="{rec.departement_id}">{rec.departement_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Employee</label>
                    <div class="input-group mb-3">
                        <input type="text" 
                            bind:value="{nmemployee_field}" 
                            disabled
                            class="form-control" placeholder="Employee" >
                        <Button on:click={() => {
                                ShowEmployee();
                            }} 
                            button_function="New"
                            button_title="<i class='bi bi-search'></i>"
                            button_css="btn-warning"/>
                    </div>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Branch</label>
                    <select
                        bind:value="{idbranch_field}" 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listBranch as rec}
                        <option value="{rec.branch_id}">{rec.branch_name}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Currency</label>
                    <select
                        bind:value="{idcurr_field}" 
                        class="required form-control ">
                        <option value="">--Please Select--</option>
                        {#each listCurr as rec}
                        <option value="{rec.curr_id}">{rec.curr_id}</option>
                        {/each}
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Tipe Document</label>
                    <select
                        class="form-control required"
                        bind:value={tipedoc_field}>
                        <option value="">--Please Select--</option>
                        <option value="ITEM">ITEM</option>
                        <option value="SERVICE">SERVICE</option>
                    </select>
                </div>
                <div class="mb-3">
                    <label for="exampleForm" class="form-label">Remark</label>
                    <textarea 
                        style="height: 100px;resize: none;" 
                        bind:value={remark_field} class="form-control "/>
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
            <div class="{page_right}">
                <div class="table-responsive border border-primary p-2" style="height: 550px;">
                    <i on:click={() => {
                        pageFullScreen();
                    }} class="bi bi-list" style="cursor: pointer;"></i>
                    {#if lock_document}
                    <div class="float-end">
                        <Button on:click={() => {
                                ShowFormDetail();
                            }} 
                            button_function=""
                            button_title="<i class='bi bi-plus-lg'></i>&nbsp;New Item / Service"
                            button_css="btn-dark"/>
                    </div>
                    {/if}
                    <table class="table table-sm">
                        <thead>
                            <tr>
                                {#if lock_document}
                                <th NOWRAP width="1%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">&nbsp;</th>
                                {/if}
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DESCRIPTION</th>
                                <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">QTY</th>
                                <th NOWRAP width="15%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UOM</th>
                                <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PRICE</th>
                                <th NOWRAP width="15%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">TOTAL</th>
                            </tr>
                        </thead>
                        <tbody>
                        {#each listdetail_field as rec}
                            <tr>
                                {#if lock_document}
                                    <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font}; cursor:pointer;">
                                        <i on:click={() => {
                                            handleDeleteListDetail(rec.detail_iditem);
                                        }} class="bi bi-trash"></i>
                                    </td>
                                {/if}
                                <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                                    {rec.detail_iditem +"-"+ rec.detail_nmitem}
                                    {#if rec.detail_purpose != ""}
                                        <br />
                                        PURPOSE : {@html rec.detail_purpose}
                                    {/if}
                                </td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.detail_qty)}</td>
                                <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.detail_iduom}</td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.detail_price)}</td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};">{new Intl.NumberFormat().format(rec.detail_total)}</td>
                            </tr>
                        {/each}
                        </tbody>
                    </table>
                </div>
                <div class="float-end">
                    Subtotal : {new Intl.NumberFormat().format(subtotal_detail)}
                </div>
                
               
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
            {#if lock_document}
            <Button on:click={() => {
                    handleSave();
                }} 
                button_function="SAVE"
                button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
                button_css="btn-warning"/>
            {/if}
        {/if}
        {#if status_field == "OPEN"}
            <Button on:click={() => {
                    handleStatusPurchaseRequest("CANCEL");
                }} 
                button_function="PROCESS"
                button_title="<i class='bi bi-trash'></i>&nbsp;&nbsp;Cancel"
                button_css="btn-danger"/>
            <Button on:click={() => {
                    handleStatusPurchaseRequest("PROCESS");
                }} 
                button_function="PROCESS"
                button_title="<i class='bi bi-arrow-right'></i>&nbsp;&nbsp;Process"
                button_css="btn-info"/>
        {/if}
	</slot:template>
</Modal>

<Modal
  modal_id="modalemployee"
  modal_size="modal-dialog-centered"
  modal_title="EMPLOYEE"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={false}>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
          <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
          <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NIK</th>
          <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">EMPLOYEE</th>
        </tr>
      </thead>
      <tbody>
        {#each listEmployee as rec}
          <tr style="cursor: pointer;" on:click={() => {
                handle_pilihemployee(rec.employee_id,rec.employee_name);
            }} >
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.employee_no}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.employee_id}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.employee_name}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
</Modal>

<Modal
	modal_id="modalcruddetail"
	modal_size="modal-dialog-centered "
	modal_title="Detail/New"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Item</label>
            <div class="input-group mb-3">
                <input type="text" 
                    bind:value={nmitemdisplay_item_field}
                    disabled
                    class="form-control" placeholder="Item" >
                <Button on:click={() => {
                        ShowItem();
                    }} 
                    button_function="New"
                    button_title="<i class='bi bi-search'></i>"
                    button_css="btn-warning"/>
            </div>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Uom</label>
            <select
                bind:value="{iduom_item_field}" 
                class="required form-control ">
                <option value="">--Please Select--</option>
                {#each listuom as rec}
                <option value="{rec.itemuom_iduom}">{rec.itemuom_iduom}</option>
                {/each}
            </select>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Qty</label>
            <Input_custom
                bind:value={qty_item_field}
                input_tipe="number_standart"
                input_required="required"
                input_maxlength="10"
                input_placeholder="Qty"/>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">price</label>
            <Input_custom
                bind:value={price_item_field}
                input_tipe="number_standart"
                input_required="required"
                input_maxlength="10"
                input_placeholder="Price"/>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Purpose</label>
            <textarea 
                style="height: 100px;resize: none;" 
                bind:value={purpose_item_field} class="form-control "/>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleListDetail();
            }} 
            button_function="SAVE"
            button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
	modal_id="modalcruddetailservice"
	modal_size="modal-dialog-centered "
	modal_title="Detail/New"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Description</label>
            <textarea 
                style="height: 100px;resize: none;" 
                bind:value={desc_item_field} class="form-control required"/>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Uom</label>
            <select
                bind:value="{iduom_item_field}" 
                class="required form-control ">
                <option value="">--Please Select--</option>
                {#each listuom as rec}
                <option value="{rec.itemuom_iduom}">{rec.itemuom_iduom}</option>
                {/each}
            </select>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Qty</label>
            <Input_custom
                bind:value={qty_item_field}
                input_tipe="number_standart"
                input_required="required"
                input_maxlength="10"
                input_placeholder="Qty"/>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">price</label>
            <Input_custom
                bind:value={price_item_field}
                input_tipe="number_standart"
                input_required="required"
                input_maxlength="10"
                input_placeholder="Price"/>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Purpose</label>
            <textarea 
                style="height: 100px;resize: none;" 
                bind:value={purpose_item_field} class="form-control "/>
        </div>
	</slot:template>
	<slot:template slot="footer">
        {#if flag_btnsave}
        <Button on:click={() => {
                handleListDetail();
            }} 
            button_function="SAVE"
            button_title="<i class='bi bi-save'></i>&nbsp;&nbsp;Save"
            button_css="btn-warning"/>
        {/if}
	</slot:template>
</Modal>

<Modal
  modal_id="modalitem"
  modal_size="modal-dialog-centered modal-lg"
  modal_title="ITEM"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={false}>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
          <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CODE</th>
          <th width="15%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CATEGORY</th>
          <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">ITEM</th>
        </tr>
      </thead>
      <tbody>
        {#each listitem as rec}
          <tr style="cursor:pointer;" on:click={() => {
                handle_pilihitem(rec.itemshare_id,rec.itemshare_name,rec.itemshare_uom);
            }}>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};text-decoration:underline;color:blue;">{rec.itemshare_id}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.itemshare_nmcateitem}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.itemshare_name}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
</Modal>