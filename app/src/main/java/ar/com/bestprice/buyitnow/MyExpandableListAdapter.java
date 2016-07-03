package ar.com.bestprice.buyitnow;

import android.app.Activity;
import android.util.SparseArray;
import android.view.LayoutInflater;
import android.view.View;
import android.view.View.OnClickListener;
import android.view.ViewGroup;
import android.widget.BaseExpandableListAdapter;
import android.widget.CheckedTextView;
import android.widget.TextView;
import android.widget.Toast;

import ar.com.bestprice.buyitnow.dto.Item;


public class MyExpandableListAdapter extends BaseExpandableListAdapter {

    private final SparseArray<Group> groups;
    public LayoutInflater inflater;
    public Activity activity;

    public MyExpandableListAdapter(Activity act, SparseArray<Group> groups) {
        activity = act;
        this.groups = groups;
        inflater = act.getLayoutInflater();
    }

    @Override
    public Object getChild(int groupPosition, int childPosition) {

        return groups.get(groupPosition).children.get(childPosition);
    }

    @Override
    public long getChildId(int groupPosition, int childPosition) {
        return 0;
    }

    @Override
    public View getChildView(int groupPosition, final int childPosition,
                             boolean isLastChild, View convertView, ViewGroup parent) {

        final Item children = (Item) getChild(groupPosition, childPosition);

        if (convertView == null) {
            convertView = inflater.inflate(R.layout.listrow_details, null);
        }


        TextView text = (TextView) convertView.findViewById(R.id.listrow_item_description);
        text.setText(children.getDescription());

        Category category = children.getCategory();

        int icon = R.drawable.wallet_icon;
        if (children.getCategory() != null) {
                icon = children.getCategory().getIcon();
        }

        text.setCompoundDrawablesWithIntrinsicBounds(icon, 0, 0, 0);

        text = (TextView) convertView.findViewById(R.id.item_price);


        text.setText(String.format("%.2f", children.getPrice()) + " $");

        convertView.setOnClickListener(new OnClickListener() {
            @Override
            public void onClick(View v) {
                Toast.makeText(activity, "ID:" + children.getId(), Toast.LENGTH_SHORT).show();
            }
        });
        return convertView;
    }

    @Override
    public int getChildrenCount(int groupPosition) {
        return groups.get(groupPosition).children.size();

    }

    @Override
    public Object getGroup(int groupPosition) {
        return groups.get(groupPosition);
    }

    @Override
    public int getGroupCount() {
        return groups.size();
    }

    @Override
    public void onGroupCollapsed(int groupPosition) {
        super.onGroupCollapsed(groupPosition);
    }

    @Override
    public void onGroupExpanded(int groupPosition) {
        super.onGroupExpanded(groupPosition);
    }

    @Override
    public long getGroupId(int groupPosition) {
        return 0;
    }

    @Override
    public View getGroupView(int groupPosition, boolean isExpanded,
                             View convertView, ViewGroup parent) {

        if (convertView == null) {
            convertView = inflater.inflate(R.layout.listrow_group, null);
        }

        Group group = (Group) getGroup(groupPosition);

        float total = 0;
        for(Item item:group.children){
            total += item.getPrice();
        }
        if(total >= 1000) {
            ((CheckedTextView) convertView).setCheckMarkDrawable(R.drawable.ic_trending_up_black_18dp);
        }else {
            ((CheckedTextView) convertView).setCheckMarkDrawable(R.drawable.ic_trending_down_black_18dp);
        }
        ((CheckedTextView) convertView).setText(group.string + "\t\t" + "$" + total);
        ((CheckedTextView) convertView).setChecked(isExpanded);
        return convertView;
    }

    @Override
    public boolean hasStableIds() {
        return false;
    }

    @Override
    public boolean isChildSelectable(int groupPosition, int childPosition) {
        return false;
    }
}