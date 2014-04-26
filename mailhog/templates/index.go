package templates

func Index() string {
	return `
<style>
  .messages {
    padding-top: 51px;
    height: 30%;
    min-height: 280px;
    border-bottom: 1px solid #999;
  }
  #messages-container {
    height: 100%;
  }
  .preview #headers {
    border-bottom: 1px solid #DDDDDD;
  }
  .selected {
    background: #0066CC;
    color: #fff;
  }
  table td, table th {
    padding: 2px 4px 2px 8px !important;
  }
  #messages-container table tbody {
    overflow: auto;
    height: 263px;
  }
  table thead {
    background: #eee;
  }
  #messages-container > table > thead > tr, #messages-container > table > tbody {
    display: block;
  }
  table th:not(:last-child) {
    border-right: 2px solid #DDD;
  }
  table#headers {
    margin-bottom: 1px;
    background: #eee;
  }
  #content .nav>li>a {
    padding: 5px 8px;
  }
  .preview #headers tbody td {
    width: 100%;
  }
  .preview #headers tbody th {
    white-space: nowrap;
    padding-right: 10px !important;
    padding-left: 10px !important;
    text-align: right;
    color: #666;
    font-weight: normal;
  }
  #preview-plain, #preview-source {
    white-space: pre;
    font-family: Courier New, Courier, System, fixed-width;
  }
  .preview .tab-content {
    padding: 0;
    overflow-y: scroll;
  }
</style>
<script>
  var columns = [15,15,40,20,10];
  var reflow = function() {
    if($('.preview').length > 0) {
      var remaining = $(window).height() - $('.preview .nav-tabs').offset().top
      $('.preview .tab-content').height(remaining - 32)
      $('#messages-container table').height($('#messages-container').height())
      $('#messages-container table tbody').height($('#messages-container').height() - $('#messages-container table thead').height())
    }

    var $table = $('#messages-container table'), $bodyCells = $table.find('tbody tr:first').children(), colWidth;
    var colWidth = [];
    for(var i in columns) {
      colWidth[i] = $table.innerWidth() / 100 * columns[i];
    }
    //var spare = $table.width() - $table.find('tbody tr:first').width();
    //var add = spare / cols;
    //colWidth = $bodyCells.map(function() { return $(this).width(); }).get();
    $table.find('thead tr').children().each(function(i, v) { $(v).width(colWidth[i]); }); 
    $table.find('tbody tr:first').children().each(function(i, v) { $(v).width(colWidth[i]); }); 
  }
  $(function() {
    reflow();
    var reflowPending = false;
    $("#messages-container table tbody").bind("DOMSubtreeModified", function() {
      if(!reflowPending) {
        reflowPending = true
        window.setTimeout(function() {
          reflow();
          reflowPending = false;
        }, 100);
      }
    });
    $(window).resize(function() {
        reflow();
    }).resize();
  })
</script>
<div class="modal fade" id="confirm-delete-all">
  <div class="modal-dialog">
    <div class="modal-content">
      <div class="modal-header">
        <button type="button" class="close" data-dismiss="modal" aria-hidden="true">&times;</button>
        <h4 class="modal-title">Delete all messages?</h4>
      </div>
      <div class="modal-body">
        <p>Are you sure you want to delete all messages?</p>
      </div>
      <div class="modal-footer">
        <button type="button" class="btn btn-default" data-dismiss="modal">Cancel</button>
        <button type="button" class="btn btn-danger" ng-click="deleteAllConfirm()">Delete all messages</button>
      </div>
    </div>
  </div>
</div>

<div class="messages">
  <div id="messages-container">
    <table class="table">
      <thead>
        <tr>
          <th>From</th>
          <th>To</th>
          <th>Subject</th>
          <th>Received</th>
          <th>Actions</th>
        </tr>
      </thead>
      <tbody>
        <tr ng-repeat="message in messages" ng-click="selectMessage(message)" ng-class="{ selected: message.Id == preview.Id }">
          <td>
            {{ message.From.Mailbox }}@{{ message.From.Domain }}
          </td>
          <td>
            <span ng-repeat="to in message.To">
              {{ to.Mailbox }}@{{ to.Domain }}
            </span>
          </td>
          <td>
            {{ message.Content.Headers.Subject.0 }}
          </td>
          <td>
            {{ date(message.Created) }}
          </td>
          <td>
            <button class="btn btn-xs btn-default" title="Delete" ng-click="deleteOne(message)"><span class="glyphicon glyphicon-remove"></span></button>
            <a class="btn btn-xs btn-default" title="Delete" href="/api/v1/messages/{{ message.Id }}/download"><span class="glyphicon glyphicon-save"></span></a>
          </td>
        </tr>
      </tbody>
    </table>
  </div>
</div>
<div class="preview" ng-if="preview">
  <table class="table" id="headers">
    <tbody ng-if="!previewAllHeaders">
      <tr>
        <th>From</th>
        <td>{{ preview.Content.Headers["From"][0] }}</td>
      </tr>
      <tr>
        <th>Subject</th>
        <td><strong>{{ preview.Content.Headers["Subject"][0] }}</strong></td>
      </tr>
      <tr>
        <th>To</th>
        <td>
          <button id="show-headers" ng-click="toggleHeaders(true)" type="button" class="btn btn-default pull-right btn-xs">Show headers <span class="glyphicon glyphicon-chevron-down"></span></button>
          {{ preview.Content.Headers["To"].join(', ') }}
        </td>
      </tr>
    </tbody>
    <tbody ng-if="previewAllHeaders">
      <tr ng-repeat="(header, value) in preview.Content.Headers">
        <th>
          {{ header }}
        </th>
        <td>
          <button id="hide-headers" ng-if="$last" ng-click="toggleHeaders(false)" type="button" class="btn btn-default pull-right btn-xs">Hide headers <span class="glyphicon glyphicon-chevron-up"></span></button>
          <div ng-repeat="v in value">{{ v }}</div>
        </td>
      </tr>
    </tbody>
  </table>
  <div id="content">
    <ul class="nav nav-tabs">
      <li class="active"><a href="#preview-html" data-toggle="tab">HTML</a></li>
      <li><a href="#preview-plain" data-toggle="tab">Plain text</a></li>
      <li><a href="#preview-source" data-toggle="tab">Source</a></li>
    </ul>
    <div class="tab-content">
      <div class="tab-pane active" id="preview-html" ng-bind-html="preview.previewHTML"></div>
      <div class="tab-pane" id="preview-plain">{{ getMessagePlain(preview) }}</div>
      <div class="tab-pane" id="preview-source">{{ getSource(preview) }}</div>
    </div>
  </div>
</div>
`;
}