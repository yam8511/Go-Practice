<html>
    <body>
        <form action="" method="post">
          <div>
            <label>Item</label>
            <input type="text" name="item"/>
          </div>
          <div>
            <label>Price</label>
            <input type="text" name="price"/>
          </div>
          <input type="hidden" name="token" value="{{.}}" />
          <button>OK</button>
        </form>
    </body>
</html>
