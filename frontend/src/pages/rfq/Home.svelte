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
    export let listPage = [];
	export let totalrecord = 0
    let dispatch = createEventDispatcher();
	let title_page = "REQUEST FOR QUOTATION"
    let sData = "";
    let myModal_newentry = "";
    let myModal_pr = "";
    let myModal_cruddetail = "";
    let myModal_vendor = "";
    let myModal_formpo = "";
    let flag_btnsave = true;
    let lock_document = false;
    let idvendor_field = "";
    let nmvendor_field = "";
    let idbranch_field = "";
    let idcurr_field = "";
    let tipedoc_field = "";
    let listdetail_field = [];
    let totalitem_field = 0;
    let status_field = "";
    let create_field = "";
    let update_field = "";

    //===EMPLOYEE==
    let listVendor = []

    //===ITEM==
    let qty_item_field = 0
    let price_item_field = 0
    let totalitem_detail = 0
    let subtotal_detail = 0

    //PURCAHSE_REQUEST
    let listpurchaserequest = [];
    let pr_iddocument = "";
    let pr_document = "";
    let pr_departement = "";
    let pr_employee = "";
    let pr_item_iditem = "";
    let pr_item_nmitem = "";
    let pr_item_decp = "";
    let pr_item_display = "";
    let pr_item_qty = 0;
    let pr_item_uom = "";
    let pr_item_price = 0;

    //PO
    let sDataPO = "";
    let po_discount = 0;
    let po_dpp = 0;
    let po_ppn = 0.0;
    let po_ppn_persen = 0.0;
    let po_ppn_total = 0;
    let po_pph = 0.0;
    let po_pph_persen = 0.0;
    let po_pph_total = 0;
    let po_dpp_tax = 0;
    let po_grandtotal = 0;

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
    const NewData = (e,id,idbranch,idvendor,nmvendor,curr,tipedoc,status,create,update) => {
        sData = e
        if(sData == "New"){
            clearField()
            lock_document = true;
        }else{
            call_detail(id)
            lock_document = false;
            idrecord = id
            idvendor_field = idvendor;
            nmvendor_field = nmvendor;
            idbranch_field = idbranch;
            idcurr_field = curr;
            tipedoc_field = tipedoc;
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
            if(idvendor_field == ""){
                flag = false
                msg += "The Vendor is required\n"
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
            if(idvendor_field == ""){
                flag = false
                msg += "The Vendor is required\n"
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
            const res = await fetch("/api/rfqsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    rfq_search: searchHome,
                    rfq_page: parseInt(pagingnow),
                    rfq_id: idrecord,
                    rfq_idbranch: idbranch_field,
                    rfq_idvendor: idvendor_field,
                    rfq_idcurr: idcurr_field,
                    rfq_tipedoc: tipedoc_field,
                    rfq_listdetail: listdetail_field,
                    rfq_totalitem: parseFloat(totalitem_field),
                    rfq_subtotal: parseFloat(subtotal_detail),
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
    async function handleStatusRFQ(e) {
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
            const res = await fetch("/api/rfqstatussave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    rfq_id: idrecord,
                    rfq_status: e,
                }),
            });
            const json = await res.json();
            if (json.status == 200) {
                flag_btnsave = true;
                status_field = "";
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
    const handleCreatePO = () => {
        po_dpp = subtotal_detail + po_dpp
        po_dpp_tax = po_dpp + po_ppn_total + po_pph_total
        po_grandtotal = po_dpp_tax
        sDataPO = "New"
        myModal_formpo = new bootstrap.Modal(document.getElementById("modalcreatepo"));
        myModal_formpo.show();
    }
    async function handleGeneratePO() {
        // po_discount = 0;
        // po_dpp = 0;
        // po_ppn = 0.0;
        // po_ppn_total = 0;
        // po_pph = 0.0;
        // po_pph_total = 0;
        // totalitem_detail
        // po_grandtotal = 0;
        let flag = true
        let msg = ""
        if(sDataPO == "New"){
            if(idrecord == ""){
                flag = false
                msg += "The Document is required\n"
            }
            
        }
        
        if(flag){
            flag_btnsave = false;
            css_loader = "display: inline-block;";
            msgloader = "Sending...";
            totalitem_field = listdetail_field.length
            const res = await fetch("/api/rfqsave", {
                method: "POST",
                headers: {
                    "Content-Type": "application/json",
                    Authorization: "Bearer " + token,
                },
                body: JSON.stringify({
                    sdata: sData,
                    page:"CURR-SAVE",
                    rfq_search: searchHome,
                    rfq_page: parseInt(pagingnow),
                    rfq_id: idrecord,
                    rfq_idbranch: idbranch_field,
                    rfq_idvendor: idvendor_field,
                    rfq_idcurr: idcurr_field,
                    rfq_tipedoc: tipedoc_field,
                    rfq_listdetail: listdetail_field,
                    rfq_totalitem: parseFloat(totalitem_field),
                    rfq_subtotal: parseFloat(subtotal_detail),
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
    function clearField(){
        idrecord = "";
        lock_document = false;
        idbranch_field = "";
        idcurr_field = "";
        tipedoc_field = "";
        listdetail_field = [];
        subtotal_detail = 0;
        status_field = "";
        create_field = "";
        update_field = "";
    }
  
    const ShowVendor = () => {
        call_vendor()
        myModal_vendor = new bootstrap.Modal(document.getElementById("modalvendor"));
        myModal_vendor.show();
    };
    const handle_pilihvendor = (e,nm) => {
        idvendor_field = e;
        nmvendor_field = nm;

        myModal_vendor.hide();
    };
   
    const handle_pilih_pr = (idprdetail,idpr,departement,employee,iditem,nmitem,descp,qty,iduom,price) => {
        pr_iddocument = idprdetail;
        pr_document = idpr;
        pr_departement = departement;
        pr_employee = employee;
        pr_item_iditem = iditem;
        pr_item_nmitem = nmitem;
        pr_item_decp = descp;
        pr_item_display = iditem+" - "+nmitem;
        pr_item_qty = qty;
        pr_item_uom = iduom;
        pr_item_price = price;
        qty_item_field = qty
        price_item_field = price

        myModal_cruddetail = new bootstrap.Modal(document.getElementById("modalcruddetail"));
        myModal_cruddetail.show();
      
    };
    const handleListDetail = () => {
        let flag = true;
        let msg = "";
        if(pr_item_iditem == "" || pr_item_nmitem == ""){
            flag = false
            msg += "The Item is required\n"
        }
        if(pr_item_uom == ""){
            flag = false
            msg += "The Uom is required\n"
        }
        if(parseFloat(qty_item_field)<1){
            flag = false
            msg += "The Qty is required\n"
        }
        if(parseFloat(qty_item_field)>parseFloat(pr_item_qty)){
            flag = false
            msg += "The Qty exceeds purchase request Qty is required\n"
        }
        if(parseFloat(price_item_field)<1){
            flag = false
            msg += "The Price is required\n"
        }
        

        for(let i=0;i<listdetail_field.length;i++){
            if(listdetail_field[i].detail_iditem == pr_item_iditem){
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
                    detail_id: pr_iddocument,
                    detail_document: pr_document,
                    detail_departement: pr_departement,
                    detail_employee: pr_employee,
                    detail_iditem: pr_item_iditem,
                    detail_nmitem: pr_item_nmitem,
                    detail_descpitem: pr_item_decp,
                    detail_qty: parseFloat(qty_item_field),
                    detail_iduom: pr_item_uom,
                    detail_price: parseFloat(price_item_field),
                    detail_total: parseFloat(total),
                },
            ];

         
            qty_item_field = 0
            price_item_field = 0
            myModal_cruddetail.hide();
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
        totalitem_detail = 0;
        for (var i = 0; i < temp.length; i++) {
            let total = parseFloat(temp[i].detail_qty)* parseFloat(temp[i].detail_price);
            subtotal_detail = subtotal_detail + total;
            totalitem_detail = totalitem_detail + 1;
            listdetail_field = [
            ...listdetail_field,
                {
                    detail_id: temp[i].detail_id,
                    detail_document: temp[i].detail_document,
                    detail_departement: temp[i].detail_departement,
                    detail_employee: temp[i].detail_employee,
                    detail_iditem: temp[i].detail_iditem,
                    detail_nmitem: temp[i].detail_nmitem,
                    detail_descpitem: temp[i].detail_descpitem,
                    detail_qty: temp[i].detail_qty,
                    detail_iduom: temp[i].detail_iduom,
                    detail_price: temp[i].detail_price,
                    detail_total: total,

                },
            ];
        }
    }
    const ShowFormPRDETAIL = () => {
        if(tipedoc_field == "" || idbranch_field == ""){
            alert("the Tipe Document and Branch is required")
        }else{
            call_purchaserequest(tipedoc_field,idbranch_field)

            myModal_pr = new bootstrap.Modal(document.getElementById("modalprdetail"));
            myModal_pr.show();
            
        }
    };
    async function call_purchaserequest(tipedoc,idbranch) {
        listpurchaserequest = [];
        const res = await fetch("/api/purchaserequestdetailview", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                purchaserequest_tipedoc: tipedoc,
                purchaserequest_idbranch: idbranch,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                let qty = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    qty = parseFloat(record[i]["prdetailview_qty"]) - parseFloat(record[i]["prdetailview_qty_po"])
                    if(qty > 0){
                        listpurchaserequest = [
                            ...listpurchaserequest,
                            {
                                prdetailview_no: no,
                                prdetailview_id: record[i]["prdetailview_id"],
                                prdetailview_idpurchaserequest: record[i]["prdetailview_idpurchaserequest"],
                                prdetailview_date: record[i]["prdetailview_date"],
                                prdetailview_iditem: record[i]["prdetailview_iditem"],
                                prdetailview_nmbranch: record[i]["prdetailview_nmbranch"],
                                prdetailview_nmdepartement: record[i]["prdetailview_nmdepartement"],
                                prdetailview_nmemployee: record[i]["prdetailview_nmemployee"],
                                prdetailview_nmitem: record[i]["prdetailview_nmitem"],
                                prdetailview_descitem: record[i]["prdetailview_descitem"],
                                prdetailview_qty: qty,
                                prdetailview_price: record[i]["prdetailview_price"],
                                prdetailview_iduom: record[i]["prdetailview_iduom"],
                            },
                        ];
                    }
                    
                }
            }
        }
    }
    async function call_vendor() {
        listVendor = [];
        const res = await fetch("/api/vendorshare", {
        method: "POST",
        headers: {
            "Content-Type": "application/json",
            Authorization: "Bearer " + token,
        },
        body: JSON.stringify({
            vendor_search: "",
            vendor_page: 0,
        }),
        });
        const json = await res.json();
        if (json.status == 200) {
        let record = json.record;
        if (record != null) {
            let no = 0;
            for (var i = 0; i < record.length; i++) {
            no = no + 1;
            listVendor = [
                ...listVendor,
                {
                    vendor_no: no,
                    vendor_id: record[i]["vendor_id"],
                    vendor_name: record[i]["vendor_name"],
                    vendor_nmcatevendor: record[i]["vendor_nmcatevendor"],
                },
            ];
            }
        }
        }
    }
    async function call_detail(idpurchase) {
        listdetail_field = [];
        subtotal_detail = 0;
        totalitem_detail = 0;
        const res = await fetch("/api/rfqdetail", {
            method: "POST",
            headers: {
                "Content-Type": "application/json",
                Authorization: "Bearer " + token,
            },
            body: JSON.stringify({
                rfq_id: idpurchase,
            }),
        });
        const json = await res.json();
        if (json.status == 200) {
            let record = json.record;
            if (record != null) {
                let no = 0;
                for (var i = 0; i < record.length; i++) {
                    no = no + 1;
                    totalitem_detail = totalitem_detail + 1;
                    let total = parseFloat(record[i]["rfqdetail_qty"]) * parseFloat(record[i]["rfqdetail_price"])
                    subtotal_detail = subtotal_detail + total;
                    listdetail_field = [
                        ...listdetail_field,
                        {
                            detail_no: no,
                            detail_id: record[i]["rfqdetail_id"],
                            detail_document: record[i]["rfqdetail_idpurchaserequest"],
                            detail_departement: record[i]["rfqdetail_nmdepartement"],
                            detail_employee: record[i]["rfqdetail_nmemployee"],
                            detail_iditem: record[i]["rfqdetail_iditem"],
                            detail_nmitem: record[i]["rfqdetail_nmitem"],
                            detail_purpose: record[i]["rfqdetail_descitem"],
                            detail_qty: record[i]["rfqdetail_qty"],
                            detail_iduom: record[i]["rfqdetail_iduom"],
                            detail_price: record[i]["rfqdetail_price"],
                            detail_total: total,
                        },
                    ];
                }
            }
        }
        console.log(totalitem_detail)
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
    const handleKeyboard_discount = (e) => {
        po_dpp = subtotal_detail - parseInt(e.target.value)
        po_ppn = 0
        po_ppn_total = 0
        po_ppn_persen = 0.0
        po_pph = 0
        po_pph_total = 0
        po_pph_persen = 0.0
        po_dpp_tax = po_dpp + po_ppn_total + po_pph_total
        po_grandtotal = po_dpp_tax
	}
    const handleKeyboard_ppn = (e) => {
        po_ppn_total = po_dpp * e.target.value
        po_ppn_persen = e.target.value * 100
        po_dpp_tax = po_dpp + po_ppn_total + po_pph_total
        po_grandtotal = po_dpp_tax
	}
    const handleKeyboard_pph = (e) => {
        po_pph_total = po_dpp * e.target.value
        po_pph_persen = e.target.value * 100
        po_dpp_tax = po_dpp + po_ppn_total + po_pph_total
        po_grandtotal = po_dpp_tax
	}
    handleKeyboard_ppn
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
                                <th NOWRAP width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size: {table_header_font};">VENDOR</th>
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
                                            //e,id,idbranch,nmvendor,curr,tipedoc,status,create,update
                                                NewData("Edit",rec.home_id,rec.home_idbranch,rec.home_idvendor,
                                                rec.home_nmvendor,rec.home_idcurr,rec.home_tipedoc,rec.home_status,
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
                                    <td  style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.home_nmvendor}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_totalitem)}</td>
                                    <td  style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.home_idcurr}</td>
                                    <td  style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.home_totalrfq)}</td>
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
            <div class="{page_left}">
                {#if sData!="New"}
                <div class="mb-2">
                    <label for="exampleForm" class="form-label">Document</label>
                    <input type="text" 
                            bind:value="{idrecord}" 
                            disabled
                            class="form-control" placeholder="Document" >
                </div>
                {/if}
                <div class="mb-2">
                    <label for="exampleForm" class="form-label">Tipe Document</label>
                    <select
                        class="form-control required"
                        bind:value={tipedoc_field}>
                        <option value="">--Please Select--</option>
                        <option value="ITEM">ITEM</option>
                        <option value="SERVICE">SERVICE</option>
                    </select>
                </div>
                <div class="mb-2">
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
                <div class="mb-2">
                    <label for="exampleForm" class="form-label">Vendor</label>
                    <div class="input-group mb-3">
                        <input type="text" 
                            bind:value="{nmvendor_field}" 
                            disabled
                            class="form-control" placeholder="Vendor" >
                        <Button on:click={() => {
                                ShowVendor();
                            }} 
                            button_function="New"
                            button_title="<i class='bi bi-search'></i>"
                            button_css="btn-warning"/>
                    </div>
                </div>
                <div class="mb-2">
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
                <div class="table-responsive border border-primary p-2" style="height: 500px;">
                    <i on:click={() => {
                        pageFullScreen();
                    }} class="bi bi-list" style="cursor: pointer;"></i>
                    {#if lock_document}
                    <div class="float-end">
                        <Button on:click={() => {
                                ShowFormPRDETAIL();
                            }} 
                            button_function=""
                            button_title="<i class='bi bi-plus-lg'></i>&nbsp;New Purchase Request"
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
                                    {rec.detail_document }<br />
                                    {rec.detail_departement } / {rec.detail_employee}<br />
                                    {rec.detail_iditem +" - "+ rec.detail_nmitem}
                                </td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_qty)}</td>
                                <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.detail_iduom}</td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_price)}</td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_total)}</td>
                            </tr>
                        {/each}
                        </tbody>
                    </table>
                </div>
                <table class="table table-sm" style="width: 100%;">
                    <tr>
                        <td width="60%" style="text-align: right;font-size: 14px;">Subtotal</td>
                        <td width="5%" style="text-align: right;font-size: 14px;">:</td>
                        <td width="*" style="text-align: right;font-size: 14px;color:blue;">{new Intl.NumberFormat().format(subtotal_detail)}</td>
                    </tr>
                </table>
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
                    handleStatusRFQ("CANCEL");
                }} 
                button_function="PROCESS"
                button_title="<i class='bi bi-trash'></i>&nbsp;&nbsp;Cancel"
                button_css="btn-danger"/>
            <Button on:click={() => {
                    handleStatusRFQ("PROCESS");
                }} 
                button_function="PROCESS"
                button_title="<i class='bi bi-arrow-right'></i>&nbsp;&nbsp;Process"
                button_css="btn-info"/>
        {/if}
        {#if status_field == "PROCESS"}
            <Button on:click={() => {
                    handleStatusRFQ("CANCEL");
                }} 
                button_function="PROCESS"
                button_title="<i class='bi bi-trash'></i>&nbsp;&nbsp;Cancel"
                button_css="btn-danger"/>
            <Button on:click={() => {
                    handleCreatePO("PROCESS");
                }} 
                button_function=""
                button_title="<i class='bi bi-arrow-right'></i>&nbsp;&nbsp;Create PO"
                button_css="btn-info"/>
        {/if}
	</slot:template>
</Modal>

<Modal
  modal_id="modalvendor"
  modal_size="modal-dialog-centered"
  modal_title="VENDOR"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={false}>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
          <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
          <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">CATEGORY </th>
          <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">VENDOR</th>
        </tr>
      </thead>
      <tbody>
        {#each listVendor as rec}
          <tr style="cursor: pointer;" on:click={() => {
                handle_pilihvendor(rec.vendor_id,rec.vendor_name);
            }} >
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.vendor_no}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.vendor_nmcatevendor}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.vendor_name}</td>
          </tr>
        {/each}
      </tbody>
    </table>
  </slot:template>
</Modal>


<Modal
  modal_id="modalprdetail"
  modal_size="modal-dialog-centered modal-xl"
  modal_title="PURCHASE REQUEST"
  modal_body_css="height:500px; overflow-y: scroll;"
  modal_footer_css="padding:5px;"
  modal_footer={false}>
  <slot:template slot="body">
    <table class="table table-sm">
      <thead>
        <tr>
          <th width="1%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">NO</th>
          <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DOCUMENT</th>
          <th width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DATE</th>
          <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">BRANCH</th>
          <th width="10%" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DEPARTEMENT / EMPLOYEE</th>
          <th width="*" style="text-align: left;vertical-align: top;font-weight:bold;font-size:{table_header_font};">DESCP</th>
          <th width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">QTY</th>
          <th width="10%" style="text-align: center;vertical-align: top;font-weight:bold;font-size:{table_header_font};">UOM</th>
          <th width="10%" style="text-align: right;vertical-align: top;font-weight:bold;font-size:{table_header_font};">PRICE</th>
        </tr>
      </thead>
      <tbody>
        {#each listpurchaserequest as rec}
          <tr style="cursor: pointer;" on:click={() => {
                //idprdetail,idpr,departement,employee,iditem,nmitem,descp,qty,iduom,price
                handle_pilih_pr(rec.prdetailview_id,rec.prdetailview_idpurchaserequest,
                    rec.prdetailview_nmdepartement,rec.prdetailview_nmemployee,
                    rec.prdetailview_iditem,rec.prdetailview_nmitem,rec.prdetailview_descitem,
                    rec.prdetailview_qty,rec.prdetailview_iduom,rec.prdetailview_price);
            }} >
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.prdetailview_no}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.prdetailview_idpurchaserequest}</td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.prdetailview_date}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.prdetailview_nmbranch}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">{rec.prdetailview_nmdepartement} / {rec.prdetailview_nmemployee}</td>
            <td NOWRAP style="text-align: left;vertical-align: top;font-size: {table_body_font};">
                {rec.prdetailview_iditem} - {rec.prdetailview_nmitem}
            </td>
            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.prdetailview_qty)}</td>
            <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.prdetailview_iduom}</td>
            <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.prdetailview_price)}</td>
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
            <div class="alert alert-light" role="alert" style="font-size: 12px;">
                {pr_document}<br />
                {pr_item_display}<br />
            </div>
        </div>
        <div class="mb-0">
            <label for="exampleForm" class="form-label">Uom</label>
            <input type="text" 
                bind:value={pr_item_uom}
                disabled
                class="form-control" placeholder="UOM" >
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
	modal_id="modalcreatepo"
	modal_size="modal-dialog-centered modal-xl"
	modal_title="Create Purchase Order"
    modal_footer_css="padding:5px;"
	modal_footer={true}>
	<slot:template slot="body">
        <div class="row">
            <div class="{page_left}">
                {#if sData!="New"}
                <div class="mb-2">
                    <label for="exampleForm" class="form-label">Document</label>
                    <input type="text" 
                            bind:value="{idrecord}" 
                            disabled
                            class="form-control" placeholder="Document" >
                </div>
                {/if}
                <div class="mb-2">
                    <label for="exampleForm" class="form-label">Tipe Document</label>
                    <select
                        class="form-control required"
                        bind:value={tipedoc_field}>
                        <option value="">--Please Select--</option>
                        <option value="ITEM">ITEM</option>
                        <option value="SERVICE">SERVICE</option>
                    </select>
                </div>
                <div class="mb-2">
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
                <div class="mb-2">
                    <label for="exampleForm" class="form-label">Vendor</label>
                    <div class="input-group mb-3">
                        <input type="text" 
                            bind:value="{nmvendor_field}" 
                            disabled
                            class="form-control" placeholder="Vendor" >
                        <Button on:click={() => {
                                ShowVendor();
                            }} 
                            button_function="New"
                            button_title="<i class='bi bi-search'></i>"
                            button_css="btn-warning"/>
                    </div>
                </div>
                <div class="mb-2">
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
                <div class="table-responsive border border-primary p-2" style="height: 500px;">
                    <i on:click={() => {
                        pageFullScreen();
                    }} class="bi bi-list" style="cursor: pointer;"></i>
                    {#if lock_document}
                    <div class="float-end">
                        <Button on:click={() => {
                                ShowFormPRDETAIL();
                            }} 
                            button_function=""
                            button_title="<i class='bi bi-plus-lg'></i>&nbsp;New Purchase Request"
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
                                    {rec.detail_document }<br />
                                    {rec.detail_departement } / {rec.detail_employee}<br />
                                    {rec.detail_iditem +" - "+ rec.detail_nmitem}
                                </td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_qty)}</td>
                                <td NOWRAP style="text-align: center;vertical-align: top;font-size: {table_body_font};">{rec.detail_iduom}</td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_price)}</td>
                                <td NOWRAP style="text-align: right;vertical-align: top;font-size: {table_body_font};color:blue;">{new Intl.NumberFormat().format(rec.detail_total)}</td>
                            </tr>
                        {/each}
                        </tbody>
                    </table>
                </div>
                <table class="table table-sm" style="width: 100%;">
                    <tr>
                        <td width="60%" style="text-align: right;font-size: 14px;">Subtotal</td>
                        <td width="1%" style="text-align: right;font-size: 14px;">:</td>
                        <td width="*" style="text-align: right;font-size: 14px;color:blue;">{new Intl.NumberFormat().format(subtotal_detail)}</td>
                    </tr>
                    <tr>
                        <td width="60%" style="text-align: right;font-size: 14px; vertical-align: top;">Discount</td>
                        <td width="1%" style="text-align: right;font-size: 14px; vertical-align: top;">:</td>
                        <td width="*" style="text-align: right;font-size: 14px; vertical-align: top;color:blue;">
                            <input 
                                bind:value={po_discount}
                                on:keyup={handleKeyboard_discount}
                                style="text-align: right;color:red;"
                                type="text">
                            <div style="margin-top: 1px;font-size: 12px;color:red;">{new Intl.NumberFormat().format(po_discount)}</div>
                        </td>
                    </tr>
                    <tr>
                        <td width="60%" style="text-align: right;font-size: 14px;vertical-align: top;">Total After Discount</td>
                        <td width="1%" style="text-align: right;font-size: 14px;vertical-align: top;">:</td>
                        <td width="*" style="text-align: right;font-size: 14px;vertical-align: top;color:blue;">{new Intl.NumberFormat().format(po_dpp)}</td>
                    </tr>
                    {#if tipedoc_field == "SERVICE"}
                        <tr>
                            <td width="60%" style="text-align: right;font-size: 14px;vertical-align: top;">
                                PPH ({po_pph_persen}%) 
                            </td>
                            <td width="1%" style="text-align: right;font-size: 14px;vertical-align: top;">:</td>
                            <td width="*" style="text-align: right;font-size: 14px;vertical-align: top;">
                                <input 
                                    bind:value={po_pph}
                                    on:keyup={handleKeyboard_pph}
                                    style="text-align: right;color:red;"
                                    type="text">
                                <div style="margin-top: 1px;font-size: 12px;color:red;">{new Intl.NumberFormat().format(po_pph_total)}</div>
                            </td>
                        </tr>
                    {/if}
                    {#if tipedoc_field == "ITEM"}
                        <tr>
                            <td width="60%" style="text-align: right;font-size: 14px;vertical-align: top;">
                                PPN ({po_ppn_persen}%) 
                            </td>
                            <td width="1%" style="text-align: right;font-size: 14px;vertical-align: top;">:</td>
                            <td width="*" style="text-align: right;font-size: 14px;vertical-align: top;">
                                <input 
                                    bind:value={po_ppn}
                                    on:keyup={handleKeyboard_ppn}
                                    style="text-align: right;color:red;"
                                    type="text">
                                <div style="margin-top: 1px;font-size: 12px;color:red;">{new Intl.NumberFormat().format(po_ppn_total)}</div>
                            </td>
                        </tr>
                    {/if}
                    <tr>
                        <td width="20%" style="text-align: right;font-size: 14px;">Grandtotal</td>
                        <td width="1%" style="text-align: right;font-size: 14px;">:</td>
                        <td width="*" style="text-align: right;font-size: 14px;color:blue;">{new Intl.NumberFormat().format(po_grandtotal)}</td>
                    </tr>
                </table>
                
               
            </div>
        </div>
	</slot:template>
	<slot:template slot="footer">
        <Button on:click={() => {
                handleGeneratePO();
            }} 
            button_function=""
            button_title="<i class='bi bi-arrow-right'></i>&nbsp;&nbsp;Generate PO"
            button_css="btn-info"/>
	</slot:template>
</Modal>
