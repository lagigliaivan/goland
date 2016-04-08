package ar.com.bestprice.buyitnow.dto;

import java.io.Serializable;

/**
 * Created by ivan on 31/03/16.
 */
public class Item implements Serializable{

    private String id;
    private String description;
    private Float price;

    public String getId() {
        return id;
    }

    public void setId(String id) {
        this.id = id;
    }

    public String getDescription() {
        return description;
    }

    public void setDescription(String description) {
        this.description = description;
    }

    public Float getPrice() {
        return price;
    }

    public void setPrice(Float price) {
        this.price = price;
    }

    @Override
    public String toString() {
        return  description + "\t \t"+ price;
    }
}
